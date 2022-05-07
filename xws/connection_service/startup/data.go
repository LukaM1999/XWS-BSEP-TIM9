package startup

import (
	"dislinkt/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var connections = []*domain.Connection{
	{
		Id:         getObjectId("62706d1b624b3da748f63fe1"),
		IssuerId:   getObjectId("62706d1b624b3da748f63fe3"),
		SubjectId:  getObjectId("62706d1b624b3da748f63fe5"),
		IsApproved: false,
		Date:       time.Now(),
	},
}

var profilesPrivacy = []*domain.ProfilePrivacy{
	{
		Id:       primitive.NewObjectID(),
		UserId:   getObjectId("62706d1b624b3da748f63fe3"),
		IsPublic: false,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
