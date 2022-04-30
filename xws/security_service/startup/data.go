package startup

import (
	"dislinkt/security_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var users = []*domain.User{
	{
		Id:       primitive.NewObjectID(),
		Username: "admin",
		Password: "admin",
	},
	{
		Id:       primitive.NewObjectID(),
		Username: "user",
		Password: "user",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}