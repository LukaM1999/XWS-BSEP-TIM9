package domain

import "context"

type JobOfferStore interface {
	GetJob(ctx context.Context, id int) (*JobOffer, error)
	GetJobs(ctx context.Context) ([]*JobOffer, error)
	GetMyJobs(ctx context.Context, profileId string) ([]*JobOffer, error)
	Delete(ctx context.Context, id int) error
	DeleteSkill(ctx context.Context, skillName string) error
	DeleteAll(ctx context.Context) error
	GetRecommendations(ctx context.Context, profileId string, skills []string) ([]*JobRecommendation, error)
	CreateJob(ctx context.Context, job *JobOffer) (*JobOffer, error)
}
