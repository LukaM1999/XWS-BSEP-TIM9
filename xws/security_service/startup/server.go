package startup

import (
	"dislinkt/common/auth"
	security "dislinkt/common/proto/security_service"
	"dislinkt/security_service/application"
	"dislinkt/security_service/domain"
	"dislinkt/security_service/infrastructure/api"
	"dislinkt/security_service/infrastructure/persistence"
	"dislinkt/security_service/startup/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
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

func accessibleRoles() map[string][]string {
	const securityServicePath = "/security.SecurityService/"
	//const profileServicePath = "/profile.ProfileService/"
	//const postServicePath = "/post.PostService/"
	//const commentServicePath = "/comment.CommentService/"
	//const reactionServicePath = "/reaction.ReactionService/"
	//const connectionServicePath = "/connection.ConnectionService/"

	return map[string][]string{
		securityServicePath + "GetAll": {"admin"},
		//profileServicePath + "Update":    {"user"},
		//commentServicePath + "Create":    {"user"},
		//commentServicePath + "Delete":    {"user"},
		//reactionServicePath + "Reaction": {"user"},
		//reactionServicePath + "Delete":   {"user"},
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	userStore := server.initUserStore(mongoClient)

	jwtManager := auth.NewJWTManager("secretKey", 30*time.Minute)

	securityService := server.initSecurityService(userStore)

	userHandler := server.initUserHandler(securityService, jwtManager)

	server.startGrpcServer(userHandler, jwtManager)
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

func (server *Server) initUserHandler(service *application.SecurityService, jwtManager *auth.JWTManager) *api.UserHandler {
	return api.NewUserHandler(service, jwtManager)
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler, jwtManager *auth.JWTManager) {
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
	security.RegisterSecurityServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
