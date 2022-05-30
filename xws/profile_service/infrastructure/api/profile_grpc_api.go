package api

import (
	"context"
	pbComment "dislinkt/common/proto/comment_service"
	pbPost "dislinkt/common/proto/post_service"
	pb "dislinkt/common/proto/profile_service"
	pbSecurity "dislinkt/common/proto/security_service"
	"dislinkt/profile_service/application"
	"dislinkt/profile_service/domain"
	"github.com/go-playground/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type ProfileHandler struct {
	pb.UnimplementedProfileServiceServer
	service        *application.ProfileService
	postClient     pbPost.PostServiceClient
	commentClient  pbComment.CommentServiceClient
	securityClient pbSecurity.SecurityServiceClient
	validate       *validator.Validate
}

func NewProfileHandler(service *application.ProfileService, postClient pbPost.PostServiceClient,
	commentClient pbComment.CommentServiceClient, securityClient pbSecurity.SecurityServiceClient) *ProfileHandler {
	return &ProfileHandler{
		service:        service,
		postClient:     postClient,
		commentClient:  commentClient,
		securityClient: securityClient,
		validate:       domain.NewProfileValidator(),
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
	if err := handler.validate.Struct(profile); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
	}
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
	if oldProfile.Username != profile.Username {
		handler.securityClient.Update(context.Background(), &pbSecurity.UpdateRequest{
			Id:       profileId,
			Username: profile.Username,
		})
	}
	return &pb.UpdateResponse{
		Profile: mapProfileToPb(profile),
	}, nil
}

func (handler *ProfileHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := handler.service.Delete(request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{}, nil
}

func (handler *ProfileHandler) GenerateToken(ctx context.Context, request *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	if ctx.Value("userId") != request.Id {
		return nil, status.Errorf(codes.Unauthenticated, "user not authenticated")
	}
	token, err := handler.service.GenerateToken(request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GenerateTokenResponse{
		Token: token,
	}, nil
}
