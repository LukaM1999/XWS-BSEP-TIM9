package startup

import (
	connection "dislinkt/common/proto/connection_service"
	"dislinkt/connection_service/application"
	"dislinkt/connection_service/domain"
	"dislinkt/connection_service/infrastructure/api"
	"dislinkt/connection_service/infrastructure/persistence"
	"dislinkt/connection_service/startup/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	QueueGroup = "connection_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	connectionStore := server.initConnectionStore(mongoClient)

	connectionService := server.initConnectionService(connectionStore)

	connectionHandler := server.initConnectionHandler(connectionService)

	server.startGrpcServer(connectionHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.ConnectionDBHost, server.config.ConnectionDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initConnectionStore(client *mongo.Client) domain.ConnectionStore {
	store := persistence.NewConnectionMongoDBStore(client)
	err := store.DeleteAll()
	if err != nil {
		return nil
	}
	for _, privacy := range profilesPrivacy {
		_, err := store.CreatePrivacy(privacy)
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, Connection := range connections {
		_, err := store.Create(Connection)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initConnectionService(store domain.ConnectionStore) *application.ConnectionService {
	return application.NewConnectionService(store)
}

func (server *Server) initConnectionHandler(service *application.ConnectionService) *api.ConnectionHandler {
	return api.NewConnectionHandler(service)
}

func (server *Server) startGrpcServer(connectionHandler *api.ConnectionHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	connection.RegisterConnectionServiceServer(grpcServer, connectionHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
