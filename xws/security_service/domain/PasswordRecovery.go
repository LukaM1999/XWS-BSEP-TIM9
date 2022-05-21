package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type PasswordRecovery struct {
	Id          primitive.ObjectID `bson:"_id"`
	Username    string             `bson:"username"`
	Token       string             `bson:"token"`
	TimeCreated time.Time          `bson:"timeCreated"`
	IsRecovered bool               `bson:"isRecovered"`
}
