package startup

import (
	"dislinkt/reaction_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var reactions = []*domain.Reaction{
	{
		Id:        getObjectId("62726d1a622a3d1748f62fe2"),
		UserId:    getObjectId("62706d1b624b3da748f63fe3"),
		PostId:    getObjectId("6210611b624b2da721f63fe3"),
		Type:      0,
		CreatedAt: time.Now(),
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
