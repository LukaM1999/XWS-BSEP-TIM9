package startup

import (
	"dislinkt/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var connections = []*domain.Connection{
	{
		Id:        getObjectId("62706d1b624b3da748f63fe1"),
		IssuerId:  getObjectId("62706d1b624b3da748f63fe3"),
		SubjectId: getObjectId("62706d1b624b3da748f63fe5"),
	},
}

var posts = []*domain.Post{
	{
		Id: getObjectId("6210611b624b2da721f63fe3"),
		Profile: domain.Profile{
			Id:        getObjectId("62706d1b624b3da748f63fe3"),
			FirstName: "Luka",
			LastName:  "Miletic",
		},
		CreatedAt: time.Time{},
		Content: domain.Content{
			Text:  "Check out my github page",
			Image: "",
			Links: nil,
		},
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
