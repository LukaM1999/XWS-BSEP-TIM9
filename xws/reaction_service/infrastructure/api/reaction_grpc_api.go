package api

import (
	"context"
	"dislinkt/common/loggers"
	pb "dislinkt/common/proto/reaction_service"
	"dislinkt/reaction_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var log = loggers.NewReactionLogger()

type ReactionHandler struct {
	pb.UnimplementedReactionServiceServer
	service *application.ReactionService
}

func NewReactionHandler(service *application.ReactionService) *ReactionHandler {
	return &ReactionHandler{
		service: service,
	}
}

func (handler *ReactionHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	Reactions, err := handler.service.Get(request.PostId)
	if err != nil {
		log.WithField("postId", request.PostId).Errorf("Cannot get reactions: %v", err)
		return nil, err
	}
	response := &pb.GetResponse{
		Reactions: []*pb.Reaction{},
	}
	for _, Reaction := range Reactions {
		current := mapReactionToPb(Reaction)
		response.Reactions = append(response.Reactions, current)
	}
	return response, nil
}

func (handler *ReactionHandler) Reaction(ctx context.Context, request *pb.ReactionRequest) (*pb.ReactionResponse, error) {
	reaction := mapPbToReaction(request.Reaction)
	reaction, err := handler.service.Reaction(reaction)
	if err != nil {
		log.Errorf("Cannot react: %v", err)
		return nil, err
	}
	log.Info("Reacted")
	return &pb.ReactionResponse{
		Reaction: mapReactionToPb(reaction),
	}, nil
}

func (handler *ReactionHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := handler.service.Delete(request.Id)
	if err != nil {
		log.Errorf("Cannot delete reaction: %v", err)
		return nil, err
	}
	log.Info("Reaction deleted")
	return &pb.DeleteResponse{}, nil
}

func (handler *ReactionHandler) DeletePostReactions(ctx context.Context,
	request *pb.DeletePostReactionsRequest) (*pb.DeletePostReactionsResponse, error) {
	postId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		log.Errorf("Cannot parse postId: %v", err)
		return nil, err
	}
	err = handler.service.DeletePostReactions(postId)
	if err != nil {
		log.Errorf("Cannot delete reactions: %v", err)
		return nil, err
	}
	log.Info("Reactions deleted")
	return &pb.DeletePostReactionsResponse{}, nil
}

func (handler *ReactionHandler) GetLogs(ctx context.Context, request *pb.GetLogsRequest) (*pb.GetLogsResponse, error) {
	logs, err := handler.service.GetLogs()
	if err != nil {
		log.Errorf("GLF")
		return nil, err
	}
	log.Info("GLD")
	return &pb.GetLogsResponse{Logs: logs}, nil
}
