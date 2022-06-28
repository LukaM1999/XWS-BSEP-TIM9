package startup

import (
	"dislinkt/connection_service/domain"
	"dislinkt/connection_service/ent"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var users = []string{
	"62706d1b624b3da748f63fe3",
	"62706d1b624b4da648f53fe3",
	"62706d1b623b3da748f63fa1",
	"55306d1b623b3da748f63fa1",
}

var connections = []*domain.Connection{
	{
		IssuerId:   "62706d1b624b4da648f53fe3",
		SubjectId:  "62706d1b623b3da748f63fa1",
		IsApproved: true,
		Date:       time.Now(),
	},
	{
		IssuerId:   "55306d1b623b3da748f63fa1",
		SubjectId:  "62706d1b623b3da748f63fa1",
		IsApproved: true,
		Date:       time.Now(),
	},
}

var blockedUsers = []*ent.BlockedUser{
	{
		CreatedAt:         time.Now(),
		IssuerPrimaryKey:  "62706d1b623b3da748f63fa1",
		SubjectPrimaryKey: "55306d1b623b3da748f63fa1",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
