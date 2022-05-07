package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Comment struct {
	Id             primitive.ObjectID `bson:"_id"`
	Content        string             `bson:"content"`
	CommentCreator CommentCreator     `bson:"commentCreator"`
	PostId         primitive.ObjectID `bson:"postId"`
	DateCreated    time.Time          `bson:"dateCreated"`
}

type CommentCreator struct {
	Id        primitive.ObjectID `bson:"_id"`
	FirstName string             `bson:"firstName"`
	LastName  string             `bson:"lastName"`
}
