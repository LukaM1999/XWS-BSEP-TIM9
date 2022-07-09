package domain

type JobOfferStore interface {
	GetJob(id int) (*JobOffer, error)
	Delete(id int) error
	DeleteSkill(skillName string) error
	DeleteAll() error
	GetRecommendations(profileId string, skills []string) ([]*JobRecommendation, error)
	CreateJob(job *JobOffer) (*JobOffer, error)
}