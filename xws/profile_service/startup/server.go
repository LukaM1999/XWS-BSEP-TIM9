package startup

import (
	"dislinkt/common/auth"
	"dislinkt/common/client"
	"dislinkt/common/loggers"
	pbComment "dislinkt/common/proto/comment_service"
	pbPost "dislinkt/common/proto/post_service"
	profile "dislinkt/common/proto/profile_service"
	pbSecurity "dislinkt/common/proto/security_service"
	saga "dislinkt/common/saga/messaging"
	"dislinkt/common/saga/messaging/nats"
	"dislinkt/profile_service/application"
	"dislinkt/profile_service/domain"
	"dislinkt/profile_service/infrastructure/api"
	"dislinkt/profile_service/infrastructure/persistence"
	"dislinkt/profile_service/startup/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

var log = loggers.NewProfileLogger()

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "profile_service"
)

func accessibleRoles() map[string]string {
	const profileServicePath = "/profile.ProfileService/"

	return map[string]string{
		profileServicePath + "GetAll":        "search:all-profiles",
		profileServicePath + "GenerateToken": "write:profile-token",
		//profileServicePath + "Get":    {"user"},
		//profileServicePath + "Create": {"user"},
		//profileServicePath + "Update": {"user"},
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	profileStore := server.initProfileStore(mongoClient)

	jwtManager := auth.NewJWTManager("secretKey", 30*time.Minute)

	commandPublisher := server.initPublisher(server.config.UpdateProfileCommandSubject)
	replySubscriber := server.initSubscriber(server.config.UpdateProfileReplySubject, QueueGroup)
	updateProfileOrchestrator := server.initUpdateProfileOrchestrator(commandPublisher, replySubscriber)

	profileService := server.initProfileService(profileStore, updateProfileOrchestrator)

	commandSubscriber := server.initSubscriber(server.config.UpdateProfileCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.UpdateProfileReplySubject)
	server.initUpdateProfileHandler(profileService, replyPublisher, commandSubscriber)

	commandSubscriber = server.initSubscriber(server.config.CreateProfileCommandSubject, QueueGroup)
	replyPublisher = server.initPublisher(server.config.CreateProfileReplySubject)
	server.initCreateProfileHandler(profileService, replyPublisher, commandSubscriber)

	postClient, err := client.NewPostClient(fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort))
	if err != nil {
		log.Fatal(err)
	}

	commentClient, err := client.NewCommentClient(fmt.Sprintf("%s:%s", server.config.CommentHost, server.config.CommentPort))
	if err != nil {
		log.Fatal(err)
	}

	securityClient, err := client.NewSecurityClient(fmt.Sprintf("%s:%s", server.config.SecurityHost, server.config.SecurityPort))
	if err != nil {
		log.Fatal(err)
	}

	profileHandler := server.initProfileHandler(profileService, postClient, commentClient, securityClient)

	server.startGrpcServer(profileHandler, jwtManager)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.ProfileDBHost, server.config.ProfileDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initProfileStore(client *mongo.Client) domain.ProfileStore {
	store := persistence.NewProfileMongoDBStore(client)
	err := store.DeleteAll()
	if err != nil {
		return nil
	}
	for _, Profile := range profiles {
		err := store.Create(Profile)
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

func (server *Server) initUpdateProfileOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *application.UpdateProfileOrchestrator {
	orchestrator, err := application.NewUpdateProfileOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initProfileService(store domain.ProfileStore, orchestrator *application.UpdateProfileOrchestrator) *application.ProfileService {
	return application.NewProfileService(store, orchestrator)
}

func (server *Server) initCreateProfileHandler(service *application.ProfileService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewCreateProfileCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initUpdateProfileHandler(service *application.ProfileService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewUpdateProfileCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initProfileHandler(service *application.ProfileService, postClient pbPost.PostServiceClient,
	commentClient pbComment.CommentServiceClient, securityClient pbSecurity.SecurityServiceClient) *api.ProfileHandler {
	return api.NewProfileHandler(service, postClient, commentClient, securityClient)
}

func (server *Server) startGrpcServer(userHandler *api.ProfileHandler, jwtManager *auth.JWTManager) {
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
	profile.RegisterProfileServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
