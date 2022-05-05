package api

import (
	"context"
	pb "dislinkt/common/proto/reaction_service"
	"dislinkt/reaction_service/application"
)

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
		return nil, err
	}
	return &pb.ReactionResponse{
		Reaction: mapReactionToPb(reaction),
	}, nil
}

func (handler *ReactionHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := handler.service.Delete(request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{}, nil
}
