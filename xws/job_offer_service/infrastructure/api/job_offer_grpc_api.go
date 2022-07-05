package api

import (
	"context"
	"dislinkt/common/loggers"
	pb "dislinkt/common/proto/job_offer_service"
	"dislinkt/job_offer_service/application"
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

func (handler *JobOfferHandler) GetJob(ctx context.Context, request *pb.GetJobRequest) (*pb.GetJobResponse, error) {
	jobOffer, err := handler.service.GetJob(int(request.Id))
	if err != nil {
		log.Errorf("Cannot get job offer: %v", err)
		return nil, err
	}
	return &pb.GetJobResponse{
		JobOffer: mapJobToPb(jobOffer),
	}, nil
}

func (handler *JobOfferHandler) CreateJob(ctx context.Context, request *pb.CreateJobRequest) (*pb.CreateJobResponse, error) {
	newJobOffer, err := handler.service.CreateJob(mapPbToJob(request.JobOffer))
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
	jobOffers, err := handler.service.GetRecommendations(request.Skills)
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
	job, err := handler.service.PromoteJob(mapPbToJob(request.JobOffer), request.Token, request.Username)
	if err != nil {
		log.Errorf("Cannot promote job: %v", err)
		return nil, err
	}
	log.Info("Job promoted")
	return &pb.PromoteJobResponse{
		JobOffer: mapJobToPb(job),
	}, nil
}
