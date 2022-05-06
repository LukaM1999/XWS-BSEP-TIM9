package api

import (
	"dislinkt/common/domain"
	pb "dislinkt/common/proto/security_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

//Function to return a pb.User from a domain.User
func mapUserToPb(user *domain.User) *pb.User {
	return &pb.User{
		Id:       user.Id.Hex(),
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	}
}

//Function to return a domain.User from a pb.User
func mapPbToUser(pbUser *pb.User) *domain.User {
	return &domain.User{
		Id:       getObjectId(pbUser.Id),
		Username: pbUser.Username,
		Password: hashPassword(pbUser.Password),
		Role:     pbUser.Role,
	}
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}
