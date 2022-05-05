package startup

import (
	"dislinkt/profile_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var profiles = []*domain.Profile{
	{
		Id:             getObjectId("62706d1b624b3da748f63fe3"),
		Username:       "admin",
		FirstName:      "Luka",
		LastName:       "Miletic",
		FullName:       "LukaMiletic",
		DateOfBirth:    time.Time{},
		PhoneNumber:    "065166161616",
		Email:          "lukam@gmail.com",
		Gender:         "male",
		Biography:      "Software Engineer",
		Education:      nil,
		WorkExperience: nil,
		Skills:         nil,
		Interests:      nil,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
