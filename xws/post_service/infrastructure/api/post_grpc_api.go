package api

import (
	"context"
	"dislinkt/common/loggers"
	pbComment "dislinkt/common/proto/comment_service"
	pb "dislinkt/common/proto/post_service"
	pbReaction "dislinkt/common/proto/reaction_service"
	"dislinkt/post_service/application"
	"dislinkt/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var log = loggers.NewPostLogger()

type PostHandler struct {
	pb.UnimplementedPostServiceServer
	service        *application.PostService
	commentClient  pbComment.CommentServiceClient
	reactionClient pbReaction.ReactionServiceClient
}

func NewPostHandler(service *application.PostService, commentClient pbComment.CommentServiceClient,
	reactionClient pbReaction.ReactionServiceClient) *PostHandler {
	return &PostHandler{
		service:        service,
		commentClient:  commentClient,
		reactionClient: reactionClient,
	}
}

func (handler *PostHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	Post, err := handler.service.Get(request.Id)
	if err != nil {
		log.WithField("postId", request.Id).Errorf("Cannot get post: %v", err)
		return nil, err
	}
	PostPb := mapPostToPb(Post)
	response := &pb.GetResponse{
		Post: PostPb,
	}
	return response, nil
}

func (handler *PostHandler) GetProfilePosts(ctx context.Context, request *pb.GetPostRequest) (*pb.GetPostsResponse, error) {
	Posts, err := handler.service.GetProfilePosts(request.ProfileId)
	if err != nil {
		log.WithField("profileId", request.ProfileId).Errorf("Cannot get profile posts: %v", err)
		return nil, err
	}
	response := &pb.GetPostsResponse{
		Posts: []*pb.Post{},
	}
	for _, Post := range Posts {
		current := mapPostToPb(Post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostHandler) GetConnectionPosts(ctx context.Context, request *pb.GetPostRequest) (*pb.GetPostsResponse, error) {
	Posts, err := handler.service.GetConnectionPosts(request.ProfileId)
	if err != nil {
		log.WithField("profileId", request.ProfileId).Errorf("Cannot get connection posts: %v", err)
		return nil, err
	}
	response := &pb.GetPostsResponse{
		Posts: []*pb.Post{},
	}
	for _, Post := range Posts {
		current := mapPostToPb(Post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	post := mapPbToPost(request.Post)
	err := handler.service.Create(post)
	if err != nil {
		log.Errorf("Cannot create post: %v", err)
		return nil, err
	}
	log.Info("Post created")
	return &pb.CreateResponse{
		Post: mapPostToPb(post),
	}, nil
}

func (handler *PostHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	id := request.Id
	post := mapPbToPost(request.Post)
	err := handler.service.Update(id, post)
	if err != nil {
		log.WithField("postId", id).Errorf("Cannot update post: %v", err)
		return nil, err
	}
	log.WithField("postId", id).Infof("Post updated")
	return &pb.UpdateResponse{
		Post: mapPostToPb(post),
	}, nil
}

func (handler *PostHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := handler.service.Delete(request.Id)
	if err != nil {
		log.WithField("postId", request.Id).Errorf("Cannot delete post: %v", err)
		return nil, err
	}
	handler.commentClient.DeletePostComments(context.TODO(), &pbComment.DeletePostCommentsRequest{PostId: request.Id})
	handler.reactionClient.DeletePostReactions(context.TODO(), &pbReaction.DeletePostReactionsRequest{PostId: request.Id})
	log.WithField("postId", request.Id).Infof("Post deleted")
	return &pb.DeleteResponse{}, nil
}

func (handler *PostHandler) CreateConnection(ctx context.Context, request *pb.CreateConnectionRequest) (*pb.CreateConnectionResponse, error) {
	connection := mapPbToConnection(request.Connection)
	err := handler.service.CreateConnection(connection)
	if err != nil {
		return nil, err
	}
	log.Info("Connection created")
	return &pb.CreateConnectionResponse{
		Connection: mapConnectionToPb(connection),
	}, nil
}

func (handler *PostHandler) DeleteConnection(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		log.Errorf("Cannot parse connectionId: %v", err)
		return nil, err
	}
	err = handler.service.DeleteConnection(id)
	if err != nil {
		log.Errorf("Cannot delete connection: %v", err)
		return nil, err
	}
	log.Info("Post deleted")
	return &pb.DeleteResponse{}, nil
}

func (handler *PostHandler) UpdateProfile(ctx context.Context, request *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	id, err := primitive.ObjectIDFromHex(request.Profile.Id)
	if err != nil {
		log.Errorf("Cannot parse profileId: %v", err)
		return nil, err
	}
	profile := &domain.Profile{
		Id:        id,
		FirstName: request.Profile.FirstName,
		LastName:  request.Profile.LastName,
	}
	err = handler.service.UpdateProfile(profile.Id, profile)
	if err != nil {
		log.WithField("profileId", profile.Id).Errorf("Cannot update profile: %v", err)
		return nil, err
	}
	log.WithField("profileId", id).Infof("Profile updated")
	return &pb.UpdateProfileResponse{
		Profile: request.Profile,
	}, nil
}

func (handler *PostHandler) CreateJob(ctx context.Context, request *pb.CreateJobRequest) (*pb.CreateJobResponse, error) {
	job, err := handler.service.CreateJob(mapPbToJob(request.JobOffer))
	if err != nil {
		log.Errorf("Cannot create job: %v", err)
		return nil, err
	}
	log.Info("Job created")
	return &pb.CreateJobResponse{
		JobOffer: mapJobToPb(job),
	}, nil
}

func (handler *PostHandler) PromoteJob(ctx context.Context, request *pb.PromoteJobRequest) (*pb.PromoteJobResponse, error) {
	job, err := handler.service.PromoteJob(mapPbToJob(request.JobOffer), request.Token, request.Username)
	if err != nil {
		log.Errorf("Cannot promote job: %v", err)
		return nil, err
	}
	log.Info("Job promoted")
	return &pb.PromoteJobResponse{
		JobOffer: mapJobToPb(job),
	}, nil
}
