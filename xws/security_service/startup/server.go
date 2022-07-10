package startup

import (
	"context"
	"dislinkt/common/auth"
	"dislinkt/common/client"
	"dislinkt/common/loggers"
	profile "dislinkt/common/proto/profile_service"
	security "dislinkt/common/proto/security_service"
	saga "dislinkt/common/saga/messaging"
	"dislinkt/common/saga/messaging/nats"
	"dislinkt/common/tracer"
	"dislinkt/security_service/application"
	"dislinkt/security_service/domain"
	"dislinkt/security_service/infrastructure/api"
	"dislinkt/security_service/infrastructure/persistence"
	"dislinkt/security_service/startup/config"
	"fmt"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	otgo "github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"net"
	"net/http"
	"time"
)

var log = loggers.NewSecurityLogger()

type Server struct {
	config *config.Config
	tracer otgo.Tracer
	closer io.Closer
}

func NewServer(config *config.Config) *Server {
	tracer, closer := tracer.Init("security-service")
	otgo.SetGlobalTracer(tracer)
	return &Server{
		config: config,
		tracer: tracer,
		closer: closer,
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

	commandPublisher := server.initPublisher(server.config.CreateProfileCommandSubject)
	replySubscriber := server.initSubscriber(server.config.CreateProfileReplySubject, QueueGroup)
	createProfileOrchestrator := server.initCreateProfileOrchestrator(commandPublisher, replySubscriber)

	securityService := server.initSecurityService(userStore, createProfileOrchestrator)

	commandSubscriber := server.initSubscriber(server.config.CreateProfileCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.CreateProfileReplySubject)
	server.initCreateProfileHandler(securityService, replyPublisher, commandSubscriber)

	commandSubscriber = server.initSubscriber(server.config.UpdateProfileCommandSubject, QueueGroup)
	replyPublisher = server.initPublisher(server.config.UpdateProfileReplySubject)
	server.initUpdateProfileHandler(securityService, replyPublisher, commandSubscriber)

	jwtManager := auth.NewJWTManager("secretKey", 30*time.Minute)

	profileClient, err := client.NewProfileClient(fmt.Sprintf("%s:%s", server.config.ProfileHost, server.config.ProfilePort))
	if err != nil {
		log.Fatalf("PCF: %v", err)
	}

	userHandler := server.initUserHandler(securityService, jwtManager, profileClient)

	server.startGrpcServer(userHandler, jwtManager)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.SecurityDBHost, server.config.SecurityDBPort)
	if err != nil {
		log.Fatalf("MGF: %v", err)
	}
	return client
}

func (server *Server) initUserStore(client *mongo.Client) domain.UserStore {
	store := persistence.NewUserMongoDBStore(client)
	err := store.DeleteAll(context.TODO())
	if err != nil {
		return nil
	}
	for _, User := range users {
		_, err := store.Register(context.TODO(), User)
		if err != nil {
			log.Fatalf("RUF: %v", err)
		}
	}
	for _, rolePermission := range rolePermissions {
		_, err := store.CreateRolePermission(context.TODO(), rolePermission)
		if err != nil {
			log.Fatal("CRF: %v", err)
		}
	}
	for _, userVerification := range userVerifications {
		_, err := store.CreateUserVerification(context.TODO(), userVerification)
		if err != nil {
			log.Fatal("CUVF: %v", err)
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

func (server *Server) initCreateProfileOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *application.CreateProfileOrchestrator {
	orchestrator, err := application.NewCreateProfileOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initSecurityService(store domain.UserStore, orchestrator *application.CreateProfileOrchestrator) *application.SecurityService {
	return application.NewSecurityService(store, orchestrator)
}

func (server *Server) initCreateProfileHandler(service *application.SecurityService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewCreateProfileCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initUpdateProfileHandler(service *application.SecurityService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewUpdateProfileCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initUserHandler(service *application.SecurityService,
	jwtManager *auth.JWTManager, profileClient profile.ProfileServiceClient) *api.UserHandler {
	return api.NewUserHandler(service, jwtManager, profileClient)
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler, jwtManager *auth.JWTManager) {
	interceptor := auth.NewAuthInterceptor(jwtManager, accessibleRoles())
	tlsCredentials, err := auth.LoadTLSServerCredentials()
	if err != nil {
		log.Fatalf("TLSF: %v", err)
	}
	serverOptions := []grpc.ServerOption{
		grpc.Creds(tlsCredentials),
		grpc.ChainUnaryInterceptor(interceptor.Unary(), grpc_prometheus.UnaryServerInterceptor),
		grpc.StreamInterceptor(interceptor.Stream()),
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(serverOptions...)
	security.RegisterSecurityServiceServer(grpcServer, userHandler)
	reflection.Register(grpcServer)
	grpc_prometheus.Register(grpcServer)
	http.Handle("/metrics", promhttp.Handler())
	//serveMux := http.NewServeMux()
	//serveMux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
	//	promhttp.Handler().ServeHTTP(w, r)
	//})
	//err = http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), serveMux)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//	return
	//}
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
