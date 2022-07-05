package startup

import (
	"dislinkt/common/auth"
	"dislinkt/common/client"
	"dislinkt/common/loggers"
	job "dislinkt/common/proto/job_offer_service"
	pbProfile "dislinkt/common/proto/profile_service"
	"dislinkt/job_offer_service/application"
	"dislinkt/job_offer_service/domain"
	"dislinkt/job_offer_service/infrastructure/api"
	"dislinkt/job_offer_service/infrastructure/persistence"
	"dislinkt/job_offer_service/startup/config"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

var log = loggers.NewJobLogger()

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "job_offer_service"
)

func accessibleRoles() map[string]string {
	const jobOfferServicePath = "/jobOffer.JobOfferService/"

	return map[string]string{
		//jobOfferServicePath + "BlockUser":   "write:block",
		//jobOfferServicePath + "UnblockUser": "write:unblock",
		//connectionServicePath + "Update": {"user"},
		//connectionServicePath + "Delete": {"user"},
	}
}

func (server *Server) Start() {

	jobOfferStore := server.initJobOfferStore()

	jwtManager := auth.NewJWTManager("secretKey", 30*time.Minute)

	profileClient, err := client.NewProfileClient(fmt.Sprintf("%s:%s", server.config.ProfileHost, server.config.ProfilePort))
	if err != nil {
		log.Fatal(err)
	}

	jobOfferService := server.initJobOfferService(jobOfferStore, profileClient)

	//commandSubscriber := server.initSubscriber(server.config.CreateProfileCommandSubject, QueueGroup)
	//replyPublisher := server.initPublisher(server.config.CreateProfileReplySubject)
	//server.initCreateProfileHandler(connectionService, replyPublisher, commandSubscriber)
	//
	//commandSubscriber = server.initSubscriber(server.config.UpdateProfileCommandSubject, QueueGroup)
	//replyPublisher = server.initPublisher(server.config.UpdateProfileReplySubject)
	//server.initUpdateProfileHandler(connectionService, replyPublisher, commandSubscriber)
	//
	jobOfferHandler := server.initJobOfferHandler(jobOfferService)

	server.startGrpcServer(jobOfferHandler, jwtManager)
}

func (server *Server) initJobOfferStore() domain.JobOfferStore {
	store := persistence.NewJobOfferPostgresStore(server.config.JobOfferDBHost, server.config.JobOfferDBPort)
	err := store.DeleteAll()
	if err != nil {
		return nil
	}
	for _, job := range jobOffers {
		_, err := store.CreateJob(job)
		if err != nil {
			return nil
		}
	}
	return store
}

//func (server *Server) initPublisher(subject string) saga.Publisher {
//	publisher, err := nats.NewNATSPublisher(
//		server.config.NatsHost, server.config.NatsPort,
//		server.config.NatsUser, server.config.NatsPass, subject)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return publisher
//}
//
//func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
//	subscriber, err := nats.NewNATSSubscriber(
//		server.config.NatsHost, server.config.NatsPort,
//		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return subscriber
//}

func (server *Server) initJobOfferService(store domain.JobOfferStore, profileClient pbProfile.ProfileServiceClient) *application.JobOfferService {
	return application.NewJobOfferService(store, profileClient)
}

func (server *Server) initJobOfferHandler(service *application.JobOfferService) *api.JobOfferHandler {
	return api.NewJobOfferHandler(service)
}

func (server *Server) startGrpcServer(jobOfferHandler *api.JobOfferHandler, jwtManager *auth.JWTManager) {
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
	job.RegisterJobOfferServiceServer(grpcServer, jobOfferHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
