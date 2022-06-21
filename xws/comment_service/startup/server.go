package startup

import (
	"dislinkt/comment_service/application"
	"dislinkt/comment_service/domain"
	"dislinkt/comment_service/infrastructure/api"
	"dislinkt/comment_service/infrastructure/persistence"
	"dislinkt/comment_service/startup/config"
	"dislinkt/common/auth"
	"dislinkt/common/loggers"
	comment "dislinkt/common/proto/comment_service"
	saga "dislinkt/common/saga/messaging"
	"dislinkt/common/saga/messaging/nats"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"net"
	"time"
)

var log = loggers.NewCommentLogger()

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "comment_service"
)

func accessibleRoles() map[string]string {
	const commentServicePath = "/comment.CommentService/"

	return map[string]string{
		//commentServicePath + "Get":    {"user"},
		//commentServicePath + "Create": {"user"},
		//commentServicePath + "Delete":    {"user"},
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	commentStore := server.initCommentStore(mongoClient)

	commentService := server.initCommentService(commentStore)

	jwtManager := auth.NewJWTManager("secretKey", 30*time.Minute)

	commandSubscriber := server.initSubscriber(server.config.UpdateProfileCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.UpdateProfileReplySubject)
	server.initUpdateProfileHandler(commentService, replyPublisher, commandSubscriber)

	commandSubscriber = server.initSubscriber(server.config.DeletePostCommandSubject, QueueGroup)
	replyPublisher = server.initPublisher(server.config.DeletePostReplySubject)
	server.initDeletePostHandler(commentService, replyPublisher, commandSubscriber)

	commentHandler := server.initCommentHandler(commentService)

	server.startGrpcServer(commentHandler, jwtManager)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.CommentDBHost, server.config.CommentDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initCommentStore(client *mongo.Client) domain.CommentStore {
	store := persistence.NewCommentMongoDBStore(client)
	err := store.DeleteAll()
	if err != nil {
		return nil
	}
	for _, Comment := range comments {
		_, err := store.Create(Comment)
		if err != nil {
			log.Fatal(err)
		}
	}
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

func (server *Server) initCommentService(store domain.CommentStore) *application.CommentService {
	return application.NewCommentService(store)
}

func (server *Server) initUpdateProfileHandler(service *application.CommentService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewUpdateProfileCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initDeletePostHandler(service *application.CommentService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewDeletePostCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initCommentHandler(service *application.CommentService) *api.CommentHandler {
	return api.NewCommentHandler(service)
}

func (server *Server) startGrpcServer(commentHandler *api.CommentHandler, jwtManager *auth.JWTManager) {
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
	comment.RegisterCommentServiceServer(grpcServer, commentHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
