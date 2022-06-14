package api

import (
	"bytes"
	"context"
	"dislinkt/common/auth"
	"dislinkt/common/domain"
	"dislinkt/common/loggers"
	pbProfile "dislinkt/common/proto/profile_service"
	pb "dislinkt/common/proto/security_service"
	"dislinkt/security_service/application"
	securityDomain "dislinkt/security_service/domain"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/pquerna/otp/totp"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"text/template"
	"time"
)

var log = loggers.NewSecurityLogger()

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
		log.WithField("username", username).Errorf("Cannot get user: %v", err)
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
		log.Errorf("Cannot get all users: %v", err)
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
	//log.Info("Registering user")
	request.User.Role = "user"
	mappedUser := mapPbToUser(request.User)
	if err := handler.validate.Struct(mappedUser); err != nil {
		log.Errorf("Invalid user: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
	}
	mappedUser.Password = HashPassword(mappedUser.Password)
	registeredUser, err := handler.service.Register(mappedUser)
	if err != nil {
		log.Errorf("Cannot register user: %v", err)
		return nil, err
	}
	logger := log.WithFields(logrus.Fields{
		"userId": registeredUser.Id.Hex(),
	})

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
		logger.Errorf("Cannot create profile: %v", err)
		handler.service.Delete(registeredUser.Id)
		return nil, err
	}
	token, err := handler.service.GenerateVerificationToken()
	if err != nil {
		logger.Errorf("Cannot generate verification token: %v", err)
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
		logger.Errorf("Cannot create user verification: %v", err)
		return nil, err
	}
	err = handler.service.SendVerificationEmail(request.GetUser().GetUsername(), request.GetEmail(), userVerification.Token)
	if err != nil {
		logger.Errorf("Cannot send verification email: %v", err)
		handler.service.Delete(registeredUser.Id)
		handler.profileClient.Delete(ctx, &pbProfile.DeleteRequest{Id: registeredUser.Id.Hex()})
		return nil, err
	}
	logger.Info("User registered")
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
		log.Errorf("Cannot parse id: %v", err)
		return nil, err
	}
	username, err := handler.service.Update(id, request.Username)
	if err != nil {
		log.WithField("id", id).Errorf("Cannot update user: %v", err)
		return nil, err
	}
	log.WithField("id", id).Infof("User updated")
	return &pb.UpdateResponse{Username: username}, nil
}

func (handler *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	//log.Info("Logging in user")
	loggerUsername := log.WithFields(logrus.Fields{
		"username": req.Username,
	})
	user, err := handler.service.Get(req.GetUsername())
	if err != nil {
		loggerUsername.Errorf("Cannot get user: %v", err)
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}
	loggerId := log.WithFields(logrus.Fields{
		"userId": user.Id.Hex(),
	})
	isVerified, err := handler.service.IsVerified(req.GetUsername())
	if err != nil {
		loggerId.Errorf("Cannot check if user is verified: %v", err)
		return nil, err
	}
	if !isVerified {
		loggerId.Errorf("User is not verified")
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}
	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		loggerId.Errorf("Incorrect username/password")
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := handler.jwtManager.Generate(user, false)
	if err != nil {
		loggerId.Errorf("Cannot generate token: %v", err)
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}
	loggerId.Info("User logged in")
	return &pb.LoginResponse{AccessToken: token}, nil
}

func (handler *UserHandler) TwoFactorAuthentication(ctx context.Context, req *pb.PasswordlessLoginRequest) (*pb.LoginResponse, error) {
	secret, err := handler.service.GetOTPSecret(req.GetUsername())
	if err != nil || secret == "" {
		log.WithField("username", req.GetUsername()).Error("Cannot get OTP secret")
		return nil, status.Errorf(codes.Internal, "No passwordless login setup: %v", err)
	}

	if !totp.Validate(req.GetOtp(), secret) {
		log.WithFields(logrus.Fields{
			"username": req.GetUsername(),
			"otp":      req.GetOtp(),
		}).Errorf("Invalid OTP")
		return nil, status.Errorf(codes.Internal, "OTP is invalid")
	}
	user, err := handler.service.Get(req.GetUsername())
	if err != nil {
		log.WithField("username", req.GetUsername()).Error("Cannot get user")
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}
	loggerId := log.WithFields(logrus.Fields{
		"userId": user.Id,
	})
	isVerified, err := handler.service.IsVerified(req.GetUsername())
	if err != nil {
		loggerId.Errorf("Cannot check if user is verified: %v", err)
		return nil, err
	}
	if !isVerified {
		loggerId.Errorf("User is not verified")
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}
	token, err := handler.jwtManager.Generate(user, true)
	if err != nil {
		loggerId.Errorf("Cannot generate token: %v", err)
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}
	loggerId.Info("User logged in")
	return &pb.LoginResponse{AccessToken: token}, nil
}

func (handler *UserHandler) SetupOTP(ctx context.Context, req *pb.SetupOTPRequest) (*pb.SetupOTPResponse, error) {
	//log.Info("Setting up OTP")
	loggerUsername := log.WithFields(logrus.Fields{
		"username": req.Username,
	})
	user, err := handler.service.Get(req.GetUsername())
	if err != nil {
		loggerUsername.Errorf("Cannot get user: %v", err)
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}
	loggerId := log.WithFields(logrus.Fields{
		"userId": user.Id,
	})

	secret, qrCode, err := handler.service.SetupOTP(req.GetUsername())
	if err != nil {
		loggerUsername.Errorf("Cannot setup OTP: %v", err)
		return nil, status.Errorf(codes.Internal, "cannot setup OTP: %v", err)
	}

	loggerId.Info("OTP setup")
	return &pb.SetupOTPResponse{
		Secret: secret,
		QrCode: qrCode,
	}, nil
}

func (handler *UserHandler) PasswordlessLogin(ctx context.Context, req *pb.PasswordlessLoginRequest) (*pb.LoginResponse, error) {
	secret, err := handler.service.GetOTPSecret(req.GetUsername())
	if err != nil || secret == "" {
		log.WithField("username", req.GetUsername()).Error("Cannot get OTP secret")
		return nil, status.Errorf(codes.Internal, "No passwordless login setup: %v", err)
	}

	if !totp.Validate(req.GetOtp(), secret) {
		log.WithFields(logrus.Fields{
			"username": req.GetUsername(),
			"otp":      req.GetOtp(),
		}).Errorf("Invalid OTP")
		return nil, status.Errorf(codes.Internal, "OTP is invalid")
	}
	user, err := handler.service.Get(req.GetUsername())
	if err != nil {
		log.WithField("username", req.GetUsername()).Error("Cannot get user")
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}
	loggerId := log.WithFields(logrus.Fields{
		"userId": user.Id,
	})
	isVerified, err := handler.service.IsVerified(req.GetUsername())
	if err != nil {
		loggerId.Errorf("Cannot check if user is verified: %v", err)
		return nil, err
	}
	if !isVerified {
		loggerId.Errorf("User is not verified")
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}
	token, err := handler.jwtManager.Generate(user, false)
	if err != nil {
		loggerId.Errorf("Cannot generate token: %v", err)
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}
	loggerId.Info("User logged in")
	return &pb.LoginResponse{AccessToken: token}, nil
}

func (handler *UserHandler) VerifyUser(ctx context.Context, req *pb.VerifyUserRequest) (*httpbody.HttpBody, error) {
	message, err := handler.service.VerifyUser(req.GetToken())
	if err != nil {
		log.WithField("token", req.GetToken()).Errorf("Cannot verify user: %v", err)
		return nil, err
	}
	t, err := template.ParseFiles("./application/verified.html")
	if err != nil {
		log.Errorf("Cannot parse template: %v", err)
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

func (handler *UserHandler) RecoverPassword(ctx context.Context, req *pb.RecoverPasswordRequest) (*pb.RecoverPasswordResponse, error) {
	token, err := handler.service.GenerateVerificationToken()
	if err != nil {
		log.Errorf("Cannot generate token: %v", err)
		return nil, err
	}
	err = handler.service.CreatePasswordRecovery(&securityDomain.PasswordRecovery{
		Id:          primitive.NewObjectID(),
		Username:    req.Username,
		Token:       token,
		TimeCreated: time.Now(),
		IsRecovered: false,
	})
	if err != nil {
		log.Errorf("Cannot create password recovery: %v", err)
		return nil, err
	}
	err = handler.service.SendRecoverPasswordEmail(req.GetEmail(), req.GetUsername(), token)
	if err != nil {
		log.WithField("email", req.GetEmail()).Errorf("Cannot send email: %v", err)
		return nil, err
	}
	log.WithField("email", req.GetEmail()).Info("Password recovrty email sent")
	return &pb.RecoverPasswordResponse{}, nil
}

func (handler *UserHandler) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error) {
	err := handler.service.UpdatePassword(req.GetToken(), req.GetPassword())
	if err != nil {
		log.WithField("token", req.GetToken()).Errorf("Cannot update password: %v", err)
		return nil, err
	}
	return &pb.UpdatePasswordResponse{}, nil
}
