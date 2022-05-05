package api

import (
	"context"
	pb "dislinkt/common/proto/profile_service"
	"dislinkt/profile_service/application"
	"strings"
)

type ProfileHandler struct {
	pb.UnimplementedProfileServiceServer
	service *application.ProfileService
}

func NewProfileHandler(service *application.ProfileService) *ProfileHandler {
	return &ProfileHandler{
		service: service,
	}
}

func (handler *ProfileHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	username := request.Username
	Profile, err := handler.service.Get(username)
	if err != nil {
		return nil, err
	}
	ProfilePb := mapProfileToPb(Profile)
	response := &pb.GetResponse{
		Profile: ProfilePb,
	}
	return response, nil
}

func (handler *ProfileHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Profiles, err := handler.service.GetAll(strings.ReplaceAll(request.Search, " ", ""))
	if err != nil {
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
	profile := mapPbToProfile(request.Profile)
	err := handler.service.Create(profile)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{
		Profile: mapProfileToPb(profile),
	}, nil
}

func (handler ProfileHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	username := request.Username
	profile := mapPbToProfile(request.Profile)
	err := handler.service.Update(username, profile)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{
		Profile: mapProfileToPb(profile),
	}, nil
}
