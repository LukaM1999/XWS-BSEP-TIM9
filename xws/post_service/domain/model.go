package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	Id        primitive.ObjectID `bson:"_id"`
	Profile   Profile            `bson:"profile"`
	CreatedAt time.Time          `bson:"createdAt"`
	Content   Content            `bson:"content"`
}

type Content struct {
	Text  string   `bson:"text"`
	Image string   `bson:"image"`
	Links []string `bson:"links"`
}

type Profile struct {
	Id        primitive.ObjectID `bson:"_id"`
	FirstName string             `bson:"firstName"`
	LastName  string             `bson:"lastName"`
}

type Connection struct {
	Id         primitive.ObjectID `bson:"_id"`
	IssuerId   primitive.ObjectID `bson:"_issuerId"`
	SubjectId  primitive.ObjectID `bson:"_subjectId"`
	IsApproved bool               `bson:"isApproved"`
}
