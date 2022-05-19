package startup

import (
	auth "dislinkt/common/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var users = []*auth.User{
	{
		Id:       getObjectId("62706d1b624b3da748f63fe3"),
		Username: "admin",
		Password: "admin",
		Role:     "admin",
	},
	{
		Id:       primitive.NewObjectID(),
		Username: "user",
		Password: "user",
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
		},
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
