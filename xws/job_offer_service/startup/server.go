package startup

import (
	"context"
	"dislinkt/common/auth"
	"dislinkt/common/loggers"
	job "dislinkt/common/proto/job_offer_service"
	saga "dislinkt/common/saga/messaging"
	"dislinkt/common/saga/messaging/nats"
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

	commandPublisher := server.initPublisher(server.config.PromoteJobCommandSubject)
	replySubscriber := server.initSubscriber(server.config.PromoteJobReplySubject, QueueGroup)
	promoteJobOrchestrator := server.initPromoteJobOrchestrator(commandPublisher, replySubscriber)

	jobOfferService := server.initJobOfferService(jobOfferStore, promoteJobOrchestrator)

	commandSubscriber := server.initSubscriber(server.config.PromoteJobCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.PromoteJobReplySubject)
	server.initPromoteJobHandler(jobOfferService, replyPublisher, commandSubscriber)

	jobOfferHandler := server.initJobOfferHandler(jobOfferService)

	server.startGrpcServer(jobOfferHandler, jwtManager)
}

func (server *Server) initJobOfferStore() domain.JobOfferStore {
	store := persistence.NewJobOfferPostgresStore(server.config.JobOfferDBHost, server.config.JobOfferDBPort)
	err := store.DeleteAll(context.TODO())
	if err != nil {
		return nil
	}
	for _, job := range jobOffers {
		_, err := store.CreateJob(context.TODO(), job)
		if err != nil {
			return nil
		}
	}
	return store
}

func (server *Server) initJobOfferService(store domain.JobOfferStore, orchestrator *application.PromoteJobOrchestrator) *application.JobOfferService {
	return application.NewJobOfferService(store, orchestrator)
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

func (server *Server) initPromoteJobOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *application.PromoteJobOrchestrator {
	orchestrator, err := application.NewPromoteJobOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initJobOfferHandler(service *application.JobOfferService) *api.JobOfferHandler {
	return api.NewJobOfferHandler(service)
}

func (server *Server) initPromoteJobHandler(service *application.JobOfferService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewPromoteJobCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
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
