package startup

import (
	"crypto/tls"
	"crypto/x509"
	"dislinkt/common/auth"
	reaction "dislinkt/common/proto/reaction_service"
	"dislinkt/reaction_service/application"
	"dislinkt/reaction_service/domain"
	"dislinkt/reaction_service/infrastructure/api"
	"dislinkt/reaction_service/infrastructure/persistence"
	"dislinkt/reaction_service/startup/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
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
	QueueGroup = "reaction_service"
)

func accessibleRoles() map[string][]string {
	const reactionServicePath = "/reaction.ReactionService/"

	return map[string][]string{
		reactionServicePath + "Reaction": {"user"},
	}
}

const (
	serverCertFile   = "../../cert/server-cert.pem"
	serverKeyFile    = "../../cert/server-key.pem"
	clientCACertFile = "../../cert/ca-cert.pem"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed client's certificate
	pemClientCA, err := ioutil.ReadFile(clientCACertFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
		//InsecureSkipVerify: true,
	}

	return credentials.NewTLS(config), nil
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	reactionStore := server.initReactionStore(mongoClient)

	jwtManager := auth.NewJWTManager("secretKey", 30*time.Minute)

	reactionService := server.initReactionService(reactionStore)

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

func (server *Server) initReactionService(store domain.ReactionStore) *application.ReactionService {
	return application.NewReactionService(store)
}

func (server *Server) initReactionHandler(service *application.ReactionService) *api.ReactionHandler {
	return api.NewReactionHandler(service)
}

func (server *Server) startGrpcServer(reactionHandler *api.ReactionHandler, jwtManager *auth.JWTManager) {
	interceptor := auth.NewAuthInterceptor(jwtManager, accessibleRoles())
	tlsCredentials, err := loadTLSCredentials()
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
