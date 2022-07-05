package api

import (
	pb "dislinkt/common/proto/job_offer_service"
	"dislinkt/job_offer_service/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapJobToPb(jobOffer *domain.JobOffer) *pb.JobOffer {
	return &pb.JobOffer{
		Id:          int64(jobOffer.Id),
		ProfileId:   jobOffer.ProfileId,
		Company:     jobOffer.Company,
		Position:    jobOffer.Position,
		Description: jobOffer.Description,
		Criteria:    jobOffer.Criteria,
		Skills:      jobOffer.Skills,
		CreatedAt:   timestamppb.New(jobOffer.CreatedAt),
	}
}

func mapPbToJob(pbJob *pb.JobOffer) *domain.JobOffer {
	return &domain.JobOffer{
		Id:          int(pbJob.Id),
		ProfileId:   pbJob.ProfileId,
		Company:     pbJob.Company,
		Position:    pbJob.Position,
		Description: pbJob.Description,
		Criteria:    pbJob.Criteria,
		Skills:      pbJob.Skills,
		CreatedAt:   pbJob.CreatedAt.AsTime(),
	}
}
