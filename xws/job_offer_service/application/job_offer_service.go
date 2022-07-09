package application

import (
	"dislinkt/job_offer_service/domain"
)

type JobOfferService struct {
	store        domain.JobOfferStore
	orchestrator *PromoteJobOrchestrator
}

func NewJobOfferService(store domain.JobOfferStore, orchestrator *PromoteJobOrchestrator) *JobOfferService {
	return &JobOfferService{
		store:        store,
		orchestrator: orchestrator,
	}
}

func (service *JobOfferService) GetJob(id int) (*domain.JobOffer, error) {
	return service.store.GetJob(id)
}

func (service *JobOfferService) CreateJob(job *domain.JobOffer) (*domain.JobOffer, error) {
	return service.store.CreateJob(job)
}

func (service *JobOfferService) PromoteJob(job *domain.JobOffer, token string, username string) (*domain.JobOffer, error) {
	err := service.orchestrator.Start(token, username, *job)
	if err != nil {
		return nil, err
	}
	return job, nil
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
