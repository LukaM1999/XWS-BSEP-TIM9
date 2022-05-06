package api

import (
	"context"
	"dislinkt/common/auth"
	pb "dislinkt/common/proto/security_service"
	"dislinkt/security_service/application"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	pb.UnimplementedSecurityServiceServer
	service    *application.SecurityService
	jwtManager *auth.JWTManager
}

func NewUserHandler(service *application.SecurityService, jwtManager *auth.JWTManager) *UserHandler {
	return &UserHandler{
		service:    service,
		jwtManager: jwtManager,
	}
}

func (handler *UserHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	username := request.Username
	User, err := handler.service.Get(username)
	if err != nil {
		return nil, err
	}
	UserPb := mapUserToPb(User)
	UserPb.Password = ""
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
		current.Password = ""
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler UserHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := mapPbToUser(request.User)
	err := handler.service.Register(user)
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return &pb.RegisterResponse{
		User: mapUserToPb(user),
	}, nil
}

func (handler *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := handler.service.Get(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := handler.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pb.LoginResponse{AccessToken: token}
	return res, nil
}
