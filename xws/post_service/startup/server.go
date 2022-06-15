package startup

import (
	"dislinkt/common/auth"
	"dislinkt/common/client"
	"dislinkt/common/loggers"
	pbComment "dislinkt/common/proto/comment_service"
	post "dislinkt/common/proto/post_service"
	pbReaction "dislinkt/common/proto/reaction_service"
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

	postService := server.initPostService(postStore)

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
	err := store.DeleteAll()
	if err != nil {
		return nil
	}
	for _, connection := range connections {
		err := store.CreateConnection(connection)
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, post := range posts {
		err := store.Create(post)
		if err != nil {
			log.Fatal(err)
		}
	}

	return store
}

func (server *Server) initPostService(store domain.PostStore) *application.PostService {
	return application.NewPostService(store)
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
