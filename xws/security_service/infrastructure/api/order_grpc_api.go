package api

import (
	"context"
	pb "dislinkt/common/proto/security_service"
	"dislinkt/security_service/application"
)

type UserHandler struct {
	pb.UnimplementedSecurityServiceServer
	service *application.SecurityService
}

func NewUserHandler(service *application.SecurityService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	username := request.Username
	User, err := handler.service.Get(username)
	if err != nil {
		return nil, err
	}
	UserPb := mapUserToPb(User)
	response := &pb.GetResponse{
		User: UserPb,
	}
	return response, nil
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Users, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}
	for _, User := range Users {
		current := mapUserToPb(User)
		response.Users = append(response.Users, current)
	}
	return response, nil
}
