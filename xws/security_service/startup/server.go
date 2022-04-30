package startup

import (
	security "dislinkt/common/proto/security_service"
	"dislinkt/security_service/application"
	"dislinkt/security_service/domain"
	"dislinkt/security_service/infrastructure/api"
	"dislinkt/security_service/infrastructure/persistence"
	"dislinkt/security_service/startup/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "security_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	userStore := server.initUserStore(mongoClient)

	securityService := server.initSecurityService(userStore)

	userHandler := server.initUserHandler(securityService)

	server.startGrpcServer(userHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.SecurityDBHost, server.config.SecurityDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initUserStore(client *mongo.Client) domain.UserStore {
	store := persistence.NewUserMongoDBStore(client)
	err := store.DeleteAll()
	if err != nil {
		return nil
	}
	for _, User := range users {
		err := store.Register(User)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initSecurityService(store domain.UserStore) *application.SecurityService {
	return application.NewSecurityService(store)
}

func (server *Server) initUserHandler(service *application.SecurityService) *api.UserHandler {
	return api.NewUserHandler(service)
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	security.RegisterSecurityServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
