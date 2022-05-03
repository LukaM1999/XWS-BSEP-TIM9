package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Reaction struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserId    primitive.ObjectID `bson:"userId"`
	PostId    primitive.ObjectID `bson:"postId"`
	Type      ReactionType       `bson:"type"`
	CreatedAt time.Time          `bson:"createdAt"`
}

type ReactionType int32

const (
	LIKE ReactionType = iota
	CELEBRATE
	SUPPORT
	LOVE
	INSIGHTFUL
	CURIOUS
	DISLIKE
)
