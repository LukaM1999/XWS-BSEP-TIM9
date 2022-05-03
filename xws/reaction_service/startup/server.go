package startup

import (
	reaction "dislinkt/common/proto/reaction_service"
	"dislinkt/reaction_service/application"
	"dislinkt/reaction_service/domain"
	"dislinkt/reaction_service/infrastructure/api"
	"dislinkt/reaction_service/infrastructure/persistence"
	"dislinkt/reaction_service/startup/config"
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
	QueueGroup = "reaction_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	reactionStore := server.initReactionStore(mongoClient)

	reactionService := server.initReactionService(reactionStore)

	reactionHandler := server.initReactionHandler(reactionService)

	server.startGrpcServer(reactionHandler)
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

func (server *Server) initReactionService(store domain.ReactionStore) *application.ReactionService {
	return application.NewReactionService(store)
}

func (server *Server) initReactionHandler(service *application.ReactionService) *api.ReactionHandler {
	return api.NewReactionHandler(service)
}

func (server *Server) startGrpcServer(reactionHandler *api.ReactionHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reaction.RegisterReactionServiceServer(grpcServer, reactionHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
