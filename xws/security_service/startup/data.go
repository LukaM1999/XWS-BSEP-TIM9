package startup

import (
	auth "dislinkt/common/domain"
	"dislinkt/security_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var users = []*auth.User{
	{
		Id:       getObjectId("62706d1b624b3da748f63fe3"),
		Username: "admin",
		Password: "$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i",
		Role:     "admin",
	},
	{
		Id:       getObjectId("62706d1b624b4da648f53fe3"),
		Username: "user",
		Password: "$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i",
		Role:     "user",
	},
	{
		Id:       getObjectId("62706d1b623b3da748f63fa1"),
		Username: "peepopog",
		Password: "$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i",
		Role:     "user",
	},
	{
		Id:       getObjectId("55306d1b623b3da748f63fa1"),
		Username: "mkisic",
		Password: "$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i",
		Role:     "user",
	},
}

var rolePermissions = []*auth.RolePermission{
	{
		Role: "admin",
		Permissions: []string{
			"read:all-users",
			"read:profile",
		},
	},
	{
		Role: "user",
		Permissions: []string{
			"read:profile",
			"search:all-profiles",
			"write:block",
			"write:unblock",
		},
	},
}

var userVerifications = []*domain.UserVerification{
	{
		Id:          getObjectId("55306d1b623b3da748f63fa1"),
		Username:    "mkisic",
		Token:       "ABCDEFGHIJ",
		TimeCreated: time.Now(),
		IsVerified:  true,
	},
	{
		Id:          getObjectId("62706d1b623b3da748f63fa1"),
		Username:    "peepopog",
		Token:       "AB123CD45E",
		TimeCreated: time.Now(),
		IsVerified:  true,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
