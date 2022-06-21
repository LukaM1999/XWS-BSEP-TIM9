package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Profile struct {
	Id             primitive.ObjectID `bson:"_id"`
	Username       string             `bson:"username" validate:"username"`
	FirstName      string             `bson:"firstName" validate:"name"`
	LastName       string             `bson:"lastName" validate:"name"`
	FullName       string             `bson:"fullName"`
	DateOfBirth    time.Time          `bson:"dateOfBirth"`
	PhoneNumber    string             `bson:"phoneNumber"`
	Email          string             `bson:"email" validate:"email"`
	Gender         string             `bson:"gender"`
	IsPrivate      bool               `bson:"isPrivate"`
	Biography      string             `bson:"biography"`
	Education      []Education        `bson:"education"`
	WorkExperience []WorkExperience   `bson:"workExperience"`
	Skills         []string           `bson:"skills"`
	Interests      []string           `bson:"interests"`
	AgentToken     string             `bson:"agentToken"`
}

type Education struct {
	School       string    `bson:"school"`
	Degree       string    `bson:"degree"`
	FieldOfStudy string    `bson:"fieldOfStudy"`
	StartDate    time.Time `bson:"startDate"`
	EndDate      time.Time `bson:"endDate"`
	Grade        string    `bson:"grade"`
	Description  string    `bson:"description"`
}

type WorkExperience struct {
	Title          string    `bson:"title"`
	Company        string    `bson:"company"`
	EmploymentType string    `bson:"employmentType"`
	Location       string    `bson:"location"`
	StartDate      time.Time `bson:"startDate"`
	EndDate        time.Time `bson:"endDate"`
}
