package startup

import (
	"context"
	"dislinkt/common/auth"
	"dislinkt/common/client"
	"dislinkt/common/loggers"
	pbComment "dislinkt/common/proto/comment_service"
	post "dislinkt/common/proto/post_service"
	pbProfile "dislinkt/common/proto/profile_service"
	pbReaction "dislinkt/common/proto/reaction_service"
	saga "dislinkt/common/saga/messaging"
	"dislinkt/common/saga/messaging/nats"
	"dislinkt/post_service/application"
	"dislinkt/post_service/domain"
	"dislinkt/post_service/infrastructure/api"
	"dislinkt/post_service/infrastructure/persistence"
	"dislinkt/post_service/startup/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

var log = loggers.NewPostLogger()

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "post_service"
)

func accessibleRoles() map[string]string {
	const postServicePath = "/post.PostService/"

	return map[string]string{
		//postServicePath + "GetAll": {"admin"},
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()

	postStore := server.initPostStore(mongoClient)

	jwtManager := auth.NewJWTManager("secretKey", 30*time.Minute)

	profileClient, err := client.NewProfileClient(fmt.Sprintf("%s:%s", server.config.ProfileHost, server.config.ProfilePort))
	if err != nil {
		log.Fatal(err)
	}

	commandPublisher := server.initPublisher(server.config.DeletePostCommandSubject)
	replySubscriber := server.initSubscriber(server.config.DeletePostReplySubject, QueueGroup)
	deletePostOrchestrator := server.initDeletePostOrchestrator(commandPublisher, replySubscriber)

	postService := server.initPostService(postStore, profileClient, deletePostOrchestrator)

	commandSubscriber := server.initSubscriber(server.config.DeletePostCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.DeletePostReplySubject)
	server.initDeletePostHandler(postService, replyPublisher, commandSubscriber)

	commandSubscriber = server.initSubscriber(server.config.UpdateProfileCommandSubject, QueueGroup)
	replyPublisher = server.initPublisher(server.config.UpdateProfileReplySubject)
	server.initUpdateProfileHandler(postService, replyPublisher, commandSubscriber)

	commentClient, err := client.NewCommentClient(fmt.Sprintf("%s:%s", server.config.CommentHost, server.config.CommentPort))
	if err != nil {
		log.Fatal(err)
	}

	reactionClient, err := client.NewReactionClient(fmt.Sprintf("%s:%s", server.config.ReactionHost, server.config.ReactionPort))
	if err != nil {
		log.Fatal(err)
	}

	postHandler := server.initPostHandler(postService, commentClient, reactionClient)

	server.startGrpcServer(postHandler, jwtManager)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.PostDBHost, server.config.PostDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initPostStore(client *mongo.Client) domain.PostStore {
	store := persistence.NewPostMongoDBStore(client)
	err := store.DeleteAll(context.TODO())
	if err != nil {
		return nil
	}
	for _, connection := range connections {
		err := store.CreateConnection(context.TODO(), connection)
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, post := range posts {
		err := store.Create(context.TODO(), post)
		if err != nil {
			log.Fatal(err)
		}
	}
	//for _, job := range jobs {
	//	err := store.CreateJob(job)
	//	if err != nil {
	//}

	return store
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initDeletePostOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *application.DeletePostOrchestrator {
	orchestrator, err := application.NewDeletePostOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initPostService(store domain.PostStore, profileClient pbProfile.ProfileServiceClient, orchestrator *application.DeletePostOrchestrator) *application.PostService {
	return application.NewPostService(store, profileClient, orchestrator)
}

func (server *Server) initDeletePostHandler(service *application.PostService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewDeletePostCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initUpdateProfileHandler(service *application.PostService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewUpdateProfileCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initPostHandler(service *application.PostService, commentClient pbComment.CommentServiceClient,
	reactionClient pbReaction.ReactionServiceClient) *api.PostHandler {
	return api.NewPostHandler(service, commentClient, reactionClient)
}

func (server *Server) startGrpcServer(postHandler *api.PostHandler, jwtManager *auth.JWTManager) {
	interceptor := auth.NewAuthInterceptor(jwtManager, accessibleRoles())
	tlsCredentials, err := auth.LoadTLSServerCredentials()
	if err != nil {
		panic("cannot load TLS credentials: %w")
	}
	serverOptions := []grpc.ServerOption{
		grpc.Creds(tlsCredentials),
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(serverOptions...)
	reflection.Register(grpcServer)
	post.RegisterPostServiceServer(grpcServer, postHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
