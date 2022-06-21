package startup

import (
	"dislinkt/common/auth"
	"dislinkt/common/loggers"
	reaction "dislinkt/common/proto/reaction_service"
	saga "dislinkt/common/saga/messaging"
	"dislinkt/common/saga/messaging/nats"
	"dislinkt/reaction_service/application"
	"dislinkt/reaction_service/domain"
	"dislinkt/reaction_service/infrastructure/api"
	"dislinkt/reaction_service/infrastructure/persistence"
	"dislinkt/reaction_service/startup/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

var log = loggers.NewReactionLogger()

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "reaction_service"
)

func accessibleRoles() map[string]string {
	const reactionServicePath = "/reaction.ReactionService/"

	return map[string]string{
		//reactionServicePath + "Reaction": {"user"},
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	reactionStore := server.initReactionStore(mongoClient)

	jwtManager := auth.NewJWTManager("secretKey", 30*time.Minute)

	reactionService := server.initReactionService(reactionStore)

	commandSubscriber := server.initSubscriber(server.config.DeletePostCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.DeletePostReplySubject)
	server.initDeletePostHandler(reactionService, replyPublisher, commandSubscriber)

	reactionHandler := server.initReactionHandler(reactionService)

	server.startGrpcServer(reactionHandler, jwtManager)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.ReactionDBHost, server.config.ReactionDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initReactionStore(client *mongo.Client) domain.ReactionStore {
	store := persistence.NewReactionMongoDBStore(client)
	err := store.DeleteAll()
	if err != nil {
		return nil
	}
	for _, Reaction := range reactions {
		_, err := store.Reaction(Reaction)
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

func (server *Server) initReactionService(store domain.ReactionStore) *application.ReactionService {
	return application.NewReactionService(store)
}

func (server *Server) initDeletePostHandler(service *application.ReactionService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewDeletePostCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initReactionHandler(service *application.ReactionService) *api.ReactionHandler {
	return api.NewReactionHandler(service)
}

func (server *Server) startGrpcServer(reactionHandler *api.ReactionHandler, jwtManager *auth.JWTManager) {
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
	reaction.RegisterReactionServiceServer(grpcServer, reactionHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
