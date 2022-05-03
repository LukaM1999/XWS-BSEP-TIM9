package startup

import (
	"dislinkt/comment_service/application"
	"dislinkt/comment_service/domain"
	"dislinkt/comment_service/infrastructure/api"
	"dislinkt/comment_service/infrastructure/persistence"
	"dislinkt/comment_service/startup/config"
	comment "dislinkt/common/proto/comment_service"
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
	QueueGroup = "comment_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	commentStore := server.initCommentStore(mongoClient)

	commentService := server.initCommentService(commentStore)

	commentHandler := server.initCommentHandler(commentService)

	server.startGrpcServer(commentHandler)
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

func (server *Server) initCommentService(store domain.CommentStore) *application.CommentService {
	return application.NewCommentService(store)
}

func (server *Server) initCommentHandler(service *application.CommentService) *api.CommentHandler {
	return api.NewCommentHandler(service)
}

func (server *Server) startGrpcServer(commentHandler *api.CommentHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	comment.RegisterCommentServiceServer(grpcServer, commentHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
