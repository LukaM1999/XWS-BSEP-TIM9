package api

import (
	"context"
	"dislinkt/comment_service/application"
	"dislinkt/comment_service/domain"
	pb "dislinkt/common/proto/comment_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentHandler struct {
	pb.UnimplementedCommentServiceServer
	service *application.CommentService
}

func NewCommentHandler(service *application.CommentService) *CommentHandler {
	return &CommentHandler{
		service: service,
	}
}

func (handler *CommentHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	Comments, err := handler.service.Get(request.PostId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetResponse{
		Comments: []*pb.Comment{},
	}
	for _, Comment := range Comments {
		current := mapCommentToPb(Comment)
		response.Comments = append(response.Comments, current)
	}
	return response, nil
}

func (handler *CommentHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	comment := mapPbToComment(request.Comment)
	_, err := handler.service.Create(comment)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{
		Comment: mapCommentToPb(comment),
	}, nil
}

func (handler *CommentHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := handler.service.Delete(request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{}, nil
}

func (handler *CommentHandler) UpdateCommentCreator(ctx context.Context,
	request *pb.UpdateCommentCreatorRequest) (*pb.UpdateCommentCreatorResponse, error) {
	creatorId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	err = handler.service.UpdateCommentCreator(creatorId, &domain.CommentCreator{
		Id:        creatorId,
		FirstName: request.CommentCreator.FirstName,
		LastName:  request.CommentCreator.LastName,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateCommentCreatorResponse{
		CommentCreator: request.CommentCreator,
	}, nil
}

func (handler *CommentHandler) DeletePostComments(ctx context.Context,
	request *pb.DeletePostCommentsRequest) (*pb.DeletePostCommentsResponse, error) {
	postId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		return nil, err
	}
	err = handler.service.DeletePostComments(postId)
	if err != nil {
		return nil, err
	}
	return &pb.DeletePostCommentsResponse{}, nil
}
