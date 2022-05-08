package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Connection struct {
	Id         primitive.ObjectID `bson:"_id"`
	IssuerId   primitive.ObjectID `bson:"issuerId"`
	SubjectId  primitive.ObjectID `bson:"subjectId"`
	Date       time.Time          `bson:"date"`
	IsApproved bool               `bson:"isApproved"`
}

type ProfilePrivacy struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserId    primitive.ObjectID `bson:"userId"`
	IsPrivate bool               `bson:"isPrivate"`
}
