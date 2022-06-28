package domain

import (
	"time"
)

type Connection struct {
	Id         int       `bson:"_id"`
	IssuerId   string    `bson:"issuerId"`
	SubjectId  string    `bson:"subjectId"`
	Date       time.Time `bson:"date"`
	IsApproved bool      `bson:"isApproved"`
}
