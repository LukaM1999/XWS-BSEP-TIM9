package api

import (
	"context"
	"dislinkt/common/auth"
	pbProfile "dislinkt/common/proto/profile_service"
	pb "dislinkt/common/proto/security_service"
	"dislinkt/security_service/application"
	"github.com/pquerna/otp/totp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	pb.UnimplementedSecurityServiceServer
	service       *application.SecurityService
	jwtManager    *auth.JWTManager
	profileClient pbProfile.ProfileServiceClient
}

func NewUserHandler(service *application.SecurityService,
	jwtManager *auth.JWTManager, profileClient pbProfile.ProfileServiceClient) *UserHandler {
	return &UserHandler{
		service:       service,
		jwtManager:    jwtManager,
		profileClient: profileClient,
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
	registeredUser, err := handler.service.Register(mapPbToUser(request.User))
	if err != nil {
		return nil, err
	}
	registeredUser.Password = ""
	_, err = handler.profileClient.Create(ctx, &pbProfile.CreateRequest{
		Profile: &pbProfile.Profile{
			Id:        registeredUser.Id.Hex(),
			Username:  registeredUser.Username,
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Email:     request.Email,
		},
	})
	if err != nil {
		handler.service.Delete(registeredUser.Id)
		return nil, err
	}
	return &pb.RegisterResponse{
		User: &pb.User{
			Id:       registeredUser.Id.Hex(),
			Username: registeredUser.Username,
			Role:     registeredUser.Role,
		}}, nil
}

func (handler *UserHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	username, err := handler.service.Update(id, request.Username)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{Username: username}, nil
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

	return &pb.LoginResponse{AccessToken: token}, nil
}

func (handler *UserHandler) SetupOTP(ctx context.Context, req *pb.SetupOTPRequest) (*pb.SetupOTPResponse, error) {
	_, err := handler.service.Get(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	secret, qrCode, err := handler.service.SetupOTP(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot setup OTP: %v", err)
	}

	return &pb.SetupOTPResponse{
		Secret: secret,
		QrCode: qrCode,
	}, nil
}

func (handler *UserHandler) PasswordlessLogin(ctx context.Context, req *pb.PasswordlessLoginRequest) (*pb.LoginResponse, error) {
	secret, err := handler.service.GetOTPSecret(req.GetUsername())
	if err != nil || secret == "" {
		return nil, status.Errorf(codes.Internal, "No passwordless login setup: %v", err)
	}

	if !totp.Validate(req.GetOtp(), secret) {
		return nil, status.Errorf(codes.Internal, "OTP is invalid")
	}
	user, err := handler.service.Get(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}
	token, err := handler.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	return &pb.LoginResponse{AccessToken: token}, nil
}
