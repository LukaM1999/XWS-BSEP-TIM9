package api

import (
	"bytes"
	"context"
	"dislinkt/common/auth"
	"dislinkt/common/domain"
	pbProfile "dislinkt/common/proto/profile_service"
	pb "dislinkt/common/proto/security_service"
	"dislinkt/security_service/application"
	securityDomain "dislinkt/security_service/domain"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/pquerna/otp/totp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"text/template"
	"time"
)

type UserHandler struct {
	pb.UnimplementedSecurityServiceServer
	service       *application.SecurityService
	jwtManager    *auth.JWTManager
	profileClient pbProfile.ProfileServiceClient
	validate      *validator.Validate
}

func NewUserHandler(service *application.SecurityService,
	jwtManager *auth.JWTManager, profileClient pbProfile.ProfileServiceClient) *UserHandler {
	return &UserHandler{
		service:       service,
		jwtManager:    jwtManager,
		profileClient: profileClient,
		validate:      domain.NewUserValidator(),
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
	request.User.Role = "user"
	mappedUser := mapPbToUser(request.User)
	if err := handler.validate.Struct(mappedUser); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
	}
	mappedUser.Password = HashPassword(mappedUser.Password)
	registeredUser, err := handler.service.Register(mappedUser)
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
	token, err := handler.service.GenerateVerificationToken()
	if err != nil {
		return nil, err
	}
	userVerification, err := handler.service.CreateUserVerification(&securityDomain.UserVerification{
		Id:          primitive.NewObjectID(),
		Username:    registeredUser.Username,
		Token:       token,
		TimeCreated: time.Now(),
		IsVerified:  false,
	})
	if err != nil {
		return nil, err
	}
	err = handler.service.SendVerificationEmail(request.GetUser().GetUsername(), request.GetEmail(), userVerification.Token)
	if err != nil {
		handler.service.Delete(registeredUser.Id)
		handler.profileClient.Delete(ctx, &pbProfile.DeleteRequest{Id: registeredUser.Id.Hex()})
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

	isVerified, err := handler.service.IsVerified(req.GetUsername())
	if err != nil {
		return nil, err
	}
	if !isVerified {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
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
	isVerified, err := handler.service.IsVerified(req.GetUsername())
	if err != nil {
		return nil, err
	}
	if !isVerified {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}
	token, err := handler.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	return &pb.LoginResponse{AccessToken: token}, nil
}

func (handler *UserHandler) VerifyUser(ctx context.Context, req *pb.VerifyUserRequest) (*httpbody.HttpBody, error) {
	message, err := handler.service.VerifyUser(req.GetToken())
	if err != nil {
		return nil, err
	}
	t, err := template.ParseFiles("./application/verified.html")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var body bytes.Buffer

	t.Execute(&body, struct {
		Message string
	}{
		Message: message,
	})
	return &httpbody.HttpBody{
		ContentType: "text/html",
		Data:        body.Bytes(),
	}, nil
}
