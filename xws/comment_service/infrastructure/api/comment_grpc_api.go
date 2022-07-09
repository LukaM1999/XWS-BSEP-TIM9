package api

import (
	"context"
	"dislinkt/comment_service/application"
	"dislinkt/comment_service/domain"
	"dislinkt/common/loggers"
	pb "dislinkt/common/proto/comment_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var log = loggers.NewCommentLogger()

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
		log.WithField("postId", request.PostId).Errorf("Cannot get comments: %v", err)
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
		log.Errorf("Cannot create comment: %v", err)
		return nil, err
	}
	log.Info("Comment created")
	return &pb.CreateResponse{
		Comment: mapCommentToPb(comment),
	}, nil
}

func (handler *CommentHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := handler.service.Delete(request.Id)
	if err != nil {
		log.Errorf("Cannot delete comment: %v", err)
		return nil, err
	}
	log.Info("Comment deleted")
	return &pb.DeleteResponse{}, nil
}

func (handler *CommentHandler) UpdateCommentCreator(ctx context.Context,
	request *pb.UpdateCommentCreatorRequest) (*pb.UpdateCommentCreatorResponse, error) {
	creatorId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		log.Errorf("Cannot parse creatorId: %v", err)
		return nil, err
	}
	err = handler.service.UpdateCommentCreator(creatorId, &domain.CommentCreator{
		Id:        creatorId,
		FirstName: request.CommentCreator.FirstName,
		LastName:  request.CommentCreator.LastName,
	})
	if err != nil {
		log.WithField("creatorId", creatorId).Errorf("Cannot update comment: %v", err)
		return nil, err
	}
	log.WithField("creatorId", creatorId).Infof("Comment updated")
	return &pb.UpdateCommentCreatorResponse{
		CommentCreator: request.CommentCreator,
	}, nil
}

func (handler *CommentHandler) DeletePostComments(ctx context.Context,
	request *pb.DeletePostCommentsRequest) (*pb.DeletePostCommentsResponse, error) {
	postId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		log.Errorf("Cannot parse postId: %v", err)
		return nil, err
	}
	err = handler.service.DeletePostComments(postId)
	if err != nil {
		log.Errorf("Cannot delete comments: %v", err)
		return nil, err
	}
	log.WithField("postId", postId).Infof("Comments deleted")
	return &pb.DeletePostCommentsResponse{}, nil
}

func (handler *CommentHandler) GetLogs(ctx context.Context, request *pb.GetLogsRequest) (*pb.GetLogsResponse, error) {
	logs, err := handler.service.GetLogs()
	if err != nil {
		log.Errorf("GLF")
		return nil, err
	}
	pbLogs := make([]*pb.Log, len(logs))
	for i, log := range logs {
		pbLogs[i] = &pb.Log{
			Time:        timestamppb.New(log.Time),
			Level:       log.Level,
			Service:     "Comment service",
			Msg:         log.Msg,
			FullContent: log.FullContent,
		}
	}
	log.Info("GLD")
	return &pb.GetLogsResponse{Logs: pbLogs}, nil
}
