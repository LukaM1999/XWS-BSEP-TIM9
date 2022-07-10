package startup

import (
	"context"
	"dislinkt/common/auth"
	"dislinkt/common/client"
	"dislinkt/common/loggers"
	connection "dislinkt/common/proto/connection_service"
	pbPost "dislinkt/common/proto/post_service"
	saga "dislinkt/common/saga/messaging"
	"dislinkt/common/saga/messaging/nats"
	"dislinkt/connection_service/application"
	"dislinkt/connection_service/domain"
	"dislinkt/connection_service/infrastructure/api"
	"dislinkt/connection_service/infrastructure/persistence"
	"dislinkt/connection_service/startup/config"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

var log = loggers.NewConnectionLogger()

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "connection_service"
)

func accessibleRoles() map[string]string {
	const connectionServicePath = "/connection.ConnectionService/"

	return map[string]string{
		connectionServicePath + "BlockUser":   "write:block",
		connectionServicePath + "UnblockUser": "write:unblock",
		//connectionServicePath + "Update": {"user"},
		//connectionServicePath + "Delete": {"user"},
	}
}

func (server *Server) Start() {

	connectionStore := server.initConnectionStore()

	jwtManager := auth.NewJWTManager("secretKey", 30*time.Minute)

	connectionService := server.initConnectionService(connectionStore)

	commandSubscriber := server.initSubscriber(server.config.CreateProfileCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.CreateProfileReplySubject)
	server.initCreateProfileHandler(connectionService, replyPublisher, commandSubscriber)

	commandSubscriber = server.initSubscriber(server.config.UpdateProfileCommandSubject, QueueGroup)
	replyPublisher = server.initPublisher(server.config.UpdateProfileReplySubject)
	server.initUpdateProfileHandler(connectionService, replyPublisher, commandSubscriber)

	postClient, err := client.NewPostClient(fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort))
	if err != nil {
		log.Fatal(err)
	}
	connectionHandler := server.initConnectionHandler(connectionService, postClient)

	server.startGrpcServer(connectionHandler, jwtManager)
}

func (server *Server) initConnectionStore() domain.ConnectionStore {
	store := persistence.NewConnectionPostgresStore(server.config.ConnectionDBHost, server.config.ConnectionDBPort)
	err := store.DeleteAll(context.TODO())
	if err != nil {
		return nil
	}
	for _, user := range users {
		err := store.CreateUser(context.TODO(), user)
		if err != nil {
			return nil
		}
	}
	for _, Connection := range connections {
		_, err := store.CreateConnection(context.TODO(), Connection.IssuerId, Connection.SubjectId)
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, BlockedUser := range blockedUsers {
		_, err := store.BlockUser(context.TODO(), BlockedUser.IssuerPrimaryKey, BlockedUser.SubjectPrimaryKey)
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

func (server *Server) initConnectionService(store domain.ConnectionStore) *application.ConnectionService {
	return application.NewConnectionService(store)
}

func (server *Server) initCreateProfileHandler(service *application.ConnectionService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewCreateProfileCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initUpdateProfileHandler(service *application.ConnectionService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewUpdateProfileCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initConnectionHandler(service *application.ConnectionService, postClient pbPost.PostServiceClient) *api.ConnectionHandler {
	return api.NewConnectionHandler(service, postClient)
}

func (server *Server) startGrpcServer(connectionHandler *api.ConnectionHandler, jwtManager *auth.JWTManager) {
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
	connection.RegisterConnectionServiceServer(grpcServer, connectionHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
