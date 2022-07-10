package api

import (
	"context"
	"dislinkt/common/loggers"
	pb "dislinkt/common/proto/job_offer_service"
	"dislinkt/common/tracer"
	"dislinkt/job_offer_service/application"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var log = loggers.NewJobLogger()

type JobOfferHandler struct {
	pb.UnimplementedJobOfferServiceServer
	service *application.JobOfferService
}

func NewJobOfferHandler(service *application.JobOfferService) *JobOfferHandler {
	return &JobOfferHandler{
		service: service,
	}
}

func (handler *JobOfferHandler) GetJobs(ctx context.Context, request *pb.GetJobsRequest) (*pb.GetJobsResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "GetJobs Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	jobOffers, err := handler.service.GetJobs(ctx)
	if err != nil {
		log.Errorf("Cannot get job offer: %v", err)
		return nil, err
	}

	var jobOffersPb []*pb.JobOffer
	for _, jobOffer := range jobOffers {
		jobOffersPb = append(jobOffersPb, mapJobToPb(jobOffer))
	}

	return &pb.GetJobsResponse{
		JobOffers: jobOffersPb,
	}, nil
}

func (handler *JobOfferHandler) GetMyJobs(ctx context.Context, request *pb.GetMyJobsRequest) (*pb.GetMyJobsResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "GetMyJobs Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	jobOffers, err := handler.service.GetMyJobs(ctx, request.ProfileId)
	if err != nil {
		log.Errorf("Cannot get job offer: %v", err)
		return nil, err
	}

	var jobOffersPb []*pb.JobOffer
	for _, jobOffer := range jobOffers {
		jobOffersPb = append(jobOffersPb, mapJobToPb(jobOffer))
	}

	return &pb.GetMyJobsResponse{
		JobOffers: jobOffersPb,
	}, nil
}

func (handler *JobOfferHandler) GetJob(ctx context.Context, request *pb.GetJobRequest) (*pb.GetJobResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "GetJob Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	jobOffer, err := handler.service.GetJob(ctx, int(request.Id))
	if err != nil {
		log.Errorf("Cannot get job offer: %v", err)
		return nil, err
	}
	return &pb.GetJobResponse{
		JobOffer: mapJobToPb(jobOffer),
	}, nil
}

func (handler *JobOfferHandler) CreateJob(ctx context.Context, request *pb.CreateJobRequest) (*pb.CreateJobResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CreateJob Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	newJobOffer, err := handler.service.CreateJob(ctx, mapPbToJob(request.JobOffer))
	if err != nil {
		log.Errorf("Cannot create job offer: %v", err)
		return nil, err
	}
	log.Info("Job offer created")
	return &pb.CreateJobResponse{
		JobOffer: mapJobToPb(newJobOffer),
	}, nil
}

func (handler *JobOfferHandler) GetRecommendations(ctx context.Context, request *pb.GetRecommendationsRequest) (*pb.GetRecommendationsResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "GetRecommendations Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	jobOffers, err := handler.service.GetRecommendations(ctx, request.ProfileId, request.Skills)
	if err != nil {
		log.Errorf("Cannot get job recommendations: %v", err)
		return nil, err
	}
	var recommendationsPb []*pb.JobRecommendation
	for _, jobOffer := range jobOffers {
		recommendationsPb = append(recommendationsPb, &pb.JobRecommendation{
			JobId:      jobOffer.JobId,
			SkillCount: jobOffer.SkillCount,
		})
	}

	return &pb.GetRecommendationsResponse{
		JobRecommendations: recommendationsPb,
	}, nil
}

func (handler *JobOfferHandler) PromoteJob(ctx context.Context, request *pb.PromoteJobRequest) (*pb.PromoteJobResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "PromoteJob Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	job, err := handler.service.PromoteJob(ctx, mapPbToJob(request.JobOffer), request.Token, request.Username)
	if err != nil {
		log.Errorf("Cannot promote job: %v", err)
		return nil, err
	}
	log.Info("Job promoted")
	return &pb.PromoteJobResponse{
		JobOffer: mapJobToPb(job),
	}, nil
}

func (handler *JobOfferHandler) GetLogs(ctx context.Context, request *pb.GetLogsRequest) (*pb.GetLogsResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "GetLogs Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	logs, err := handler.service.GetLogs(ctx)
	if err != nil {
		log.Errorf("GLF")
		return nil, err
	}
	pbLogs := make([]*pb.Log, len(logs))
	for i, log := range logs {
		pbLogs[i] = &pb.Log{
			Time:        timestamppb.New(log.Time),
			Level:       log.Level,
			Service:     "Job offer service",
			Msg:         log.Msg,
			FullContent: log.FullContent,
		}
	}
	log.Info("GLD")
	return &pb.GetLogsResponse{Logs: pbLogs}, nil
}
