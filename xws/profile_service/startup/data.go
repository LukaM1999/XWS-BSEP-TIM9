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
		Education:      make([]domain.Education, 0),
		WorkExperience: make([]domain.WorkExperience, 0),
		Skills:         make([]string, 0),
		Interests:      make([]string, 0),
		AgentToken:     "",
		IsPrivate:      false,
	},
	{
		Id:             getObjectId("62706d1b623b3da748f63fa1"),
		Username:       "peepopog",
		FirstName:      "Peepo",
		LastName:       "Pog",
		FullName:       "PeepoPog",
		DateOfBirth:    time.Time{},
		PhoneNumber:    "063123321",
		Email:          "peepopog@gmail.com",
		Gender:         "male",
		Biography:      "Game Developer",
		Education:      make([]domain.Education, 0),
		WorkExperience: make([]domain.WorkExperience, 0),
		Skills:         make([]string, 0),
		Interests:      make([]string, 0),
		AgentToken:     "",
		IsPrivate:      false,
	},
	{
		Id:             getObjectId("55306d1b623b3da748f63fa1"),
		Username:       "mkisic",
		FirstName:      "Mihajlo",
		LastName:       "Kisic",
		FullName:       "MihajloKisic",
		DateOfBirth:    time.Time{},
		PhoneNumber:    "0641112233",
		Email:          "mkisic@gmail.com",
		Gender:         "male",
		Biography:      "Software Engineer",
		Education:      make([]domain.Education, 0),
		WorkExperience: make([]domain.WorkExperience, 0),
		Skills:         make([]string, 0),
		Interests:      make([]string, 0),
		AgentToken:     "",
		IsPrivate:      false,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
