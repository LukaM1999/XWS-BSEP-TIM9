package application

import (
	"context"
	pbProfile "dislinkt/common/proto/profile_service"
	"dislinkt/job_offer_service/domain"
	"errors"
)

type JobOfferService struct {
	store         domain.JobOfferStore
	profileClient pbProfile.ProfileServiceClient
}

func NewJobOfferService(store domain.JobOfferStore, profileClient pbProfile.ProfileServiceClient) *JobOfferService {
	return &JobOfferService{
		store:         store,
		profileClient: profileClient,
	}
}

func (service *JobOfferService) GetJob(id int) (*domain.JobOffer, error) {
	return service.store.GetJob(id)
}

func (service *JobOfferService) CreateJob(job *domain.JobOffer) (*domain.JobOffer, error) {
	return service.store.CreateJob(job)
}

func (service *JobOfferService) PromoteJob(job *domain.JobOffer, token string, username string) (*domain.JobOffer, error) {
	profile, err := service.profileClient.GetByToken(context.TODO(), &pbProfile.GetByTokenRequest{Token: &token})
	if err != nil {
		return nil, err
	}
	if *profile.Profile.Username != username {
		return nil, errors.New("invalid username or token")
	}
	return service.store.CreateJob(job)
}

func (service *JobOfferService) Delete(id int) error {
	return service.store.Delete(id)
}

func (service *JobOfferService) DeleteSkill(skillName string) error {
	return service.store.DeleteSkill(skillName)
}

func (service *JobOfferService) GetRecommendations(skills []string) ([]*domain.JobRecommendation, error) {
	return service.store.GetRecommendations(skills)
}
