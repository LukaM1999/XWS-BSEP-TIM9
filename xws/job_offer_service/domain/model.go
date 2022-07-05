package domain

import (
	"time"
)

type JobOffer struct {
	Id          int       `bson:"id"`
	ProfileId   string    `bson:"profileId"`
	Company     string    `bson:"company"`
	Position    string    `bson:"position"`
	Description string    `bson:"description"`
	Criteria    string    `bson:"criteria"`
	Skills      []string  `bson:"skills"`
	CreatedAt   time.Time `bson:"createdAt"`
}

type JobRecommendation struct {
	JobId      int64 `bson:"jobId"`
	SkillCount int64 `bson:"skillCount"`
}
