package startup

import (
	"dislinkt/common/auth"
	"dislinkt/common/client"
	profile "dislinkt/common/proto/profile_service"
	security "dislinkt/common/proto/security_service"
	"dislinkt/security_service/application"
	"dislinkt/security_service/domain"
	"dislinkt/security_service/infrastructure/api"
	"dislinkt/security_service/infrastructure/persistence"
	"dislinkt/security_service/startup/config"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

func accessibleRoles() map[string]string {
	const securityServicePath = "/security.SecurityService/"

	return map[string]string{
		securityServicePath + "GetAll": "read:all-users",
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	userStore := server.initUserStore(mongoClient)

	jwtManager := auth.NewJWTManager("secretKey", 30*time.Minute)

	securityService := server.initSecurityService(userStore)

	profileClient, err := client.NewProfileClient(fmt.Sprintf("%s:%s", server.config.ProfileHost, server.config.ProfilePort))
	if err != nil {
		log.Fatal(err)
	}

	userHandler := server.initUserHandler(securityService, jwtManager, profileClient)

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
		_, err := store.Register(User)
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, rolePermission := range rolePermissions {
		_, err := store.CreateRolePermission(rolePermission)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initSecurityService(store domain.UserStore) *application.SecurityService {
	return application.NewSecurityService(store)
}

func (server *Server) initUserHandler(service *application.SecurityService,
	jwtManager *auth.JWTManager, profileClient profile.ProfileServiceClient) *api.UserHandler {
	return api.NewUserHandler(service, jwtManager, profileClient)
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler, jwtManager *auth.JWTManager) {
	interceptor := auth.NewAuthInterceptor(jwtManager, accessibleRoles())
	tlsCredentials, err := auth.LoadTLSServerCredentials()
	if err != nil {
		log.Fatalf("failed to load TLS credentials: %v", err)
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
