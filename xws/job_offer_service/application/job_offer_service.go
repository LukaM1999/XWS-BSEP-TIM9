package application

import (
	auth "dislinkt/common/domain"
	"dislinkt/job_offer_service/domain"
	"os"
	"regexp"
	"strings"
	"time"
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

func (service *JobOfferService) GetRecommendations(profileId string, skills []string) ([]*domain.JobRecommendation, error) {
	return service.store.GetRecommendations(profileId, skills)
}

func (service *JobOfferService) GetLogs() ([]auth.Log, error) {
	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	content, err := os.ReadFile(logPathPrefix + "job_offer_service/job_offer.log")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	logs := make([]auth.Log, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		var log auth.Log
		splitBySpace := strings.Split(line, " ")
		log.Time, err = time.Parse("2006-01-02T15:04:05.000Z", strings.Trim(strings.Split(splitBySpace[0], "=")[1], "\""))
		if err != nil {
			log.Time = time.Time{}
		}
		log.Level = strings.Split(splitBySpace[1], "=")[1]
		re := regexp.MustCompile(`msg="[/\\=!?'"\.a-zA-Z0-9_\s:-]*"`)
		msg := re.FindString(line)
		if msg != "" {
			log.Msg = strings.Trim(strings.Split(msg, "=")[1], "\"")
		}
		if msg == "" {
			re = regexp.MustCompile(`msg=[a-zA-Z]*`)
			msg = re.FindString(line)
			if msg != "" {
				log.Msg = strings.Split(msg, "=")[1]
			}
		}
		log.FullContent = line
		logs = append(logs, log)
	}
	return logs, nil
}

func (service *JobOfferService) GetJobs() ([]*domain.JobOffer, error) {
	return service.store.GetJobs()
}

func (service *JobOfferService) GetMyJobs(profileId string) ([]*domain.JobOffer, error) {
	return service.store.GetMyJobs(profileId)
}
