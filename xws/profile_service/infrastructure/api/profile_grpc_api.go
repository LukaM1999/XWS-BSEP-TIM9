package api

import (
	"context"
	"dislinkt/common/loggers"
	pb "dislinkt/common/proto/profile_service"
	"dislinkt/profile_service/application"
	"dislinkt/profile_service/domain"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var log = loggers.NewProfileLogger()

type ProfileHandler struct {
	pb.UnimplementedProfileServiceServer
	service  *application.ProfileService
	validate *validator.Validate
}

func NewProfileHandler(service *application.ProfileService) *ProfileHandler {
	return &ProfileHandler{
		service:  service,
		validate: domain.NewProfileValidator(),
	}
}

func (handler *ProfileHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	profileId := request.Id
	Profile, err := handler.service.Get(*profileId)
	if err != nil {
		log.WithField("profileId", profileId).Errorf("Cannot get profile: %v", err)
		return nil, err
	}
	ProfilePb := mapProfileToPb(Profile)
	response := &pb.GetResponse{
		Profile: ProfilePb,
	}
	return response, nil
}

func (handler *ProfileHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Profiles, err := handler.service.GetAll(strings.ReplaceAll(*request.Search, " ", ""))
	if err != nil {
		log.Errorf("Cannot get all profiles: %v", err)
		return nil, err
	}
	response := &pb.GetAllResponse{
		Profiles: []*pb.Profile{},
	}
	for _, Profile := range Profiles {
		current := mapProfileToPb(Profile)
		response.Profiles = append(response.Profiles, current)
	}
	return response, nil
}

func (handler ProfileHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	empty := ""
	emptyBool := false
	value := &empty
	valueBool := &emptyBool
	request.Profile.PhoneNumber = value
	request.Profile.DateOfBirth = timestamppb.New(time.Now())
	request.Profile.Gender = value
	request.Profile.IsPrivate = valueBool
	request.Profile.Biography = value
	request.Profile.AgentToken = value

	profile := mapPbToProfile(request.Profile)
	if err := handler.validate.Struct(profile); err != nil {
		log.Errorf("Validation failed: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
	}
	err := handler.service.Create(profile)
	if err != nil {
		log.Errorf("Cannot create profile: %v", err)
		return nil, err
	}
	log.Info("Profile created")
	return &pb.CreateResponse{
		Profile: mapProfileToPb(profile),
	}, nil
}

func (handler ProfileHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	profile := mapPbToProfile(request.Profile)
	err := handler.service.Update(*request.Id, profile)
	if err != nil {
		log.Errorf("Cannot update profile: %v", err)
		return nil, err
	}
	log.WithField("profileId", profile.Id.Hex()).Infof("Profile updated")
	return &pb.UpdateResponse{
		Profile: mapProfileToPb(profile),
	}, nil
}

func (handler *ProfileHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := handler.service.Delete(*request.Id)
	if err != nil {
		log.Errorf("Cannot delete profile: %v", err)
		return nil, err
	}
	log.Info("Profile deleted")
	return &pb.DeleteResponse{}, nil
}

func (handler *ProfileHandler) GenerateToken(ctx context.Context, request *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	//if ctx.Value("userId").(string) != *request.Id {
	//	return nil, status.Errorf(codes.Unauthenticated, "user not authenticated")
	//}
	token, err := handler.service.GenerateToken(*request.Id)
	if err != nil {
		log.Errorf("Cannot generate token: %v", err)
		return nil, err
	}
	log.Info("Token generated")
	return &pb.GenerateTokenResponse{
		Token: &token,
	}, nil
}

func (handler *ProfileHandler) GetByToken(ctx context.Context, request *pb.GetByTokenRequest) (*pb.GetByTokenResponse, error) {
	Profile, err := handler.service.GetByToken(*request.Token)
	if err != nil {
		log.WithField("token", request.Token).Errorf("Cannot get profile: %v", err)
		return nil, err
	}
	ProfilePb := mapProfileToPb(Profile)
	response := &pb.GetByTokenResponse{
		Profile: ProfilePb,
	}
	return response, nil
}

func (handler *ProfileHandler) GetLogs(ctx context.Context, request *pb.GetLogsRequest) (*pb.GetLogsResponse, error) {
	logs, err := handler.service.GetLogs()
	if err != nil {
		log.Errorf("GLF")
		return nil, err
	}
	log.Info("GLD")
	return &pb.GetLogsResponse{Logs: logs}, nil
}
