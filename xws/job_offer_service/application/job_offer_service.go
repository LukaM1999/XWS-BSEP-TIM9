package application

import (
	"context"
	auth "dislinkt/common/domain"
	"dislinkt/common/tracer"
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

func (service *JobOfferService) GetJob(ctx context.Context, id int) (*domain.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "GetJob Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.GetJob(ctx, id)
}

func (service *JobOfferService) CreateJob(ctx context.Context, job *domain.JobOffer) (*domain.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "CreateJob Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.CreateJob(ctx, job)
}

func (service *JobOfferService) PromoteJob(ctx context.Context, job *domain.JobOffer, token string, username string) (*domain.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "PromoteJob Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.orchestrator.Start(token, username, *job)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (service *JobOfferService) Delete(ctx context.Context, id int) error {
	span := tracer.StartSpanFromContext(ctx, "Delete Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.Delete(ctx, id)
}

func (service *JobOfferService) DeleteSkill(ctx context.Context, skillName string) error {
	span := tracer.StartSpanFromContext(ctx, "DeleteSkill Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.DeleteSkill(ctx, skillName)
}

func (service *JobOfferService) GetRecommendations(ctx context.Context, profileId string, skills []string) ([]*domain.JobRecommendation, error) {
	span := tracer.StartSpanFromContext(ctx, "GetRecommendations Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.GetRecommendations(ctx, profileId, skills)
}

func (service *JobOfferService) GetLogs(ctx context.Context) ([]auth.Log, error) {
	span := tracer.StartSpanFromContext(ctx, "GetLogs Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

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

func (service *JobOfferService) GetJobs(ctx context.Context) ([]*domain.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "GetJobs Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.GetJobs(ctx)
}

func (service *JobOfferService) GetMyJobs(ctx context.Context, profileId string) ([]*domain.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "GetMyJobs Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.GetMyJobs(ctx, profileId)
}
