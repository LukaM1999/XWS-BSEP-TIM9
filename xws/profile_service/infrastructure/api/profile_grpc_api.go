package api

import (
	"context"
	pbComment "dislinkt/common/proto/comment_service"
	pbPost "dislinkt/common/proto/post_service"
	pb "dislinkt/common/proto/profile_service"
	"dislinkt/profile_service/application"
	"strings"
)

type ProfileHandler struct {
	pb.UnimplementedProfileServiceServer
	service       *application.ProfileService
	postClient    pbPost.PostServiceClient
	commentClient pbComment.CommentServiceClient
}

func NewProfileHandler(service *application.ProfileService, postClient pbPost.PostServiceClient,
	commentClient pbComment.CommentServiceClient) *ProfileHandler {
	return &ProfileHandler{
		service:       service,
		postClient:    postClient,
		commentClient: commentClient,
	}
}

func (handler *ProfileHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	profileId := request.Id
	Profile, err := handler.service.Get(profileId)
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
	oldProfile, err := handler.service.Get(request.Profile.Id)
	if err != nil {
		return nil, err
	}
	profileId := request.Id
	profile := mapPbToProfile(request.Profile)
	err = handler.service.Update(profileId, profile)
	if err != nil {
		return nil, err
	}
	if oldProfile.FirstName != profile.FirstName || oldProfile.LastName != profile.LastName {
		_, err = handler.postClient.UpdateProfile(context.Background(), &pbPost.UpdateProfileRequest{
			Profile: &pbPost.Profile{
				Id:        request.Profile.Id,
				FirstName: profile.FirstName,
				LastName:  profile.LastName,
			},
		})
		if err != nil {
			handler.service.Update(profileId, oldProfile)
			return nil, err
		}
		_, err = handler.commentClient.UpdateCommentCreator(context.Background(), &pbComment.UpdateCommentCreatorRequest{
			Id: profileId,
			CommentCreator: &pbComment.CommentCreator{
				Id:        profileId,
				FirstName: profile.FirstName,
				LastName:  profile.LastName,
			},
		})
		if err != nil {
			handler.service.Update(profileId, oldProfile)
			return nil, err
		}
	}
	return &pb.UpdateResponse{
		Profile: mapProfileToPb(profile),
	}, nil
}
