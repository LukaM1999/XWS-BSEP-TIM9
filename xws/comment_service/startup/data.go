package startup

import (
	"dislinkt/comment_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var comments = []*domain.Comment{
	{
		Id:      getObjectId("62726d1a622a3d1748f63fe2"),
		Content: "Mnogo dobar post bate",
		CommentCreator: domain.CommentCreator{
			Id:        getObjectId("62706d1b624b3da748f63fe3"),
			FirstName: "Luka",
			LastName:  "Miletic",
		},
		PostId:      getObjectId("6210611b624b2da721f63fe3"),
		DateCreated: time.Now(),
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
