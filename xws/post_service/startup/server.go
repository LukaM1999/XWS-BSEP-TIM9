package startup

import (
	"crypto/tls"
	"crypto/x509"
	"dislinkt/common/auth"
	post "dislinkt/common/proto/post_service"
	"dislinkt/post_service/application"
	"dislinkt/post_service/domain"
	"dislinkt/post_service/infrastructure/api"
	"dislinkt/post_service/infrastructure/persistence"
	"dislinkt/post_service/startup/config"
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
	QueueGroup = "post_service"
)

func accessibleRoles() map[string][]string {
	const postServicePath = "/post.PostService/"

	return map[string][]string{
		//postServicePath + "GetAll": {"admin"},
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

	postStore := server.initPostStore(mongoClient)

	jwtManager := auth.NewJWTManager("secretKey", 30*time.Minute)

	postService := server.initPostService(postStore)

	postHandler := server.initPostHandler(postService)

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

func (server *Server) initPostHandler(service *application.PostService) *api.PostHandler {
	return api.NewPostHandler(service)
}

func (server *Server) startGrpcServer(postHandler *api.PostHandler, jwtManager *auth.JWTManager) {
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
	post.RegisterPostServiceServer(grpcServer, postHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
