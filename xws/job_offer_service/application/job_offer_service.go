package application

import (
	"dislinkt/job_offer_service/domain"
	"os"
	"strings"
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

func (service *JobOfferService) GetLogs() ([]string, error) {
	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	content, err := os.ReadFile(logPathPrefix + "job_offer_service/job_offer.log")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}
