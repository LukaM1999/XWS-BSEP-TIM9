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
		log.WithField("username", username).Errorf("GUF: %v", err)
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
		log.Errorf("AUF: %v", err)
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
		log.Errorf("IUF: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
	}
	mappedUser.Password = HashPassword(mappedUser.Password)
	registeredUser, err := handler.service.Register(mappedUser, request.FirstName, request.LastName, request.Email)
	if err != nil {
		log.Errorf("RUF: %v", err)
		return nil, err
	}
	registeredUser.Password = ""
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
		log.Errorf("PIDF: %v", err)
		return nil, err
	}
	username, err := handler.service.Update(id, request.Username)
	if err != nil {
		log.WithField("id", id).Errorf("UUF: %v", err)
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
		loggerUsername.Errorf("GUF: %v", err)
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}
	loggerId := log.WithFields(logrus.Fields{
		"userId": user.Id.Hex(),
	})
	isVerified, err := handler.service.IsVerified(req.GetUsername())
	if err != nil {
		loggerId.Errorf("IUVF: %v", err)
		return nil, err
	}
	if !isVerified {
		loggerId.Errorf("NUVF")
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}
	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		loggerId.Errorf("UPF")
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := handler.jwtManager.Generate(user, false)
	if err != nil {
		loggerId.Errorf("GJWTF: %v", err)
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}
	loggerId.Info("ULGD")
	return &pb.LoginResponse{AccessToken: token}, nil
}

func (handler *UserHandler) TwoFactorAuthentication(ctx context.Context, req *pb.PasswordlessLoginRequest) (*pb.LoginResponse, error) {
	secret, err := handler.service.GetOTPSecret(req.GetUsername())
	if err != nil || secret == "" {
		log.WithField("username", req.GetUsername()).Error("GOSF")
		return nil, status.Errorf(codes.Internal, "No passwordless login setup: %v", err)
	}

	if !totp.Validate(req.GetOtp(), secret) {
		log.WithFields(logrus.Fields{
			"username": req.GetUsername(),
			"otp":      req.GetOtp(),
		}).Errorf("OTPF")
		return nil, status.Errorf(codes.Internal, "OTP is invalid")
	}
	user, err := handler.service.Get(req.GetUsername())
	if err != nil {
		log.WithField("username", req.GetUsername()).Error("GUF")
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}
	loggerId := log.WithFields(logrus.Fields{
		"userId": user.Id,
	})
	isVerified, err := handler.service.IsVerified(req.GetUsername())
	if err != nil {
		loggerId.Errorf("IUVF")
		return nil, err
	}
	if !isVerified {
		loggerId.Errorf("NUVF")
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}
	token, err := handler.jwtManager.Generate(user, true)
	if err != nil {
		loggerId.Errorf("GJWTF")
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}
	loggerId.Info("ULGD")
	return &pb.LoginResponse{AccessToken: token}, nil
}

func (handler *UserHandler) SetupOTP(ctx context.Context, req *pb.SetupOTPRequest) (*pb.SetupOTPResponse, error) {
	//log.Info("Setting up OTP")
	loggerUsername := log.WithFields(logrus.Fields{
		"username": req.Username,
	})
	user, err := handler.service.Get(req.GetUsername())
	if err != nil {
		loggerUsername.Errorf("GUF")
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}
	loggerId := log.WithFields(logrus.Fields{
		"userId": user.Id,
	})

	secret, qrCode, err := handler.service.SetupOTP(req.GetUsername())
	if err != nil {
		loggerUsername.Errorf("SOTPF")
		return nil, status.Errorf(codes.Internal, "cannot setup OTP: %v", err)
	}

	loggerId.Info("OTPS")
	return &pb.SetupOTPResponse{
		Secret: secret,
		QrCode: qrCode,
	}, nil
}

func (handler *UserHandler) PasswordlessLogin(ctx context.Context, req *pb.PasswordlessLoginRequest) (*pb.LoginResponse, error) {
	secret, err := handler.service.GetOTPSecret(req.GetUsername())
	if err != nil || secret == "" {
		log.WithField("username", req.GetUsername()).Error("GOSF")
		return nil, status.Errorf(codes.Internal, "No passwordless login setup: %v", err)
	}

	if !totp.Validate(req.GetOtp(), secret) {
		log.WithFields(logrus.Fields{
			"username": req.GetUsername(),
			"otp":      req.GetOtp(),
		}).Errorf("OTPF")
		return nil, status.Errorf(codes.Internal, "OTP is invalid")
	}
	user, err := handler.service.Get(req.GetUsername())
	if err != nil {
		log.WithField("username", req.GetUsername()).Error("GUF")
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}
	loggerId := log.WithFields(logrus.Fields{
		"userId": user.Id,
	})
	isVerified, err := handler.service.IsVerified(req.GetUsername())
	if err != nil {
		loggerId.Errorf("IUVF")
		return nil, err
	}
	if !isVerified {
		loggerId.Errorf("NUVF")
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}
	token, err := handler.jwtManager.Generate(user, false)
	if err != nil {
		loggerId.Errorf("GJWTF")
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}
	loggerId.Info("ULGD")
	return &pb.LoginResponse{AccessToken: token}, nil
}

func (handler *UserHandler) VerifyUser(ctx context.Context, req *pb.VerifyUserRequest) (*httpbody.HttpBody, error) {
	message, err := handler.service.VerifyUser(req.GetToken())
	if err != nil {
		log.WithField("token", req.GetToken()).Errorf("IUVF")
		return nil, err
	}
	t, err := template.ParseFiles("./templates/verified.html")
	if err != nil {
		log.Errorf("PTF")
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
		log.Errorf("CVTF")
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
		log.Errorf("CPWRF")
		return nil, err
	}
	err = handler.service.SendRecoverPasswordEmail(req.GetEmail(), req.GetUsername(), token)
	if err != nil {
		log.WithField("email", req.GetEmail()).Errorf("SREF")
		return nil, err
	}
	log.WithField("email", req.GetEmail()).Info("RES")
	return &pb.RecoverPasswordResponse{}, nil
}

func (handler *UserHandler) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error) {
	err := handler.service.UpdatePassword(req.GetToken(), req.GetPassword())
	if err != nil {
		log.WithField("token", req.GetToken()).Errorf("UUPF")
		return nil, err
	}
	return &pb.UpdatePasswordResponse{}, nil
}

func (handler *UserHandler) GetLogs(ctx context.Context, request *pb.GetLogsRequest) (*pb.GetLogsResponse, error) {
	logs, err := handler.service.GetLogs()
	if err != nil {
		log.Errorf("GLF")
		return nil, err
	}
	log.Info("GLD")
	return &pb.GetLogsResponse{Logs: logs}, nil
}
