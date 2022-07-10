package api

import (
	"context"
	"dislinkt/common/loggers"
	pb "dislinkt/common/proto/reaction_service"
	"dislinkt/common/tracer"
	"dislinkt/reaction_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	span := tracer.StartSpanFromContext(ctx, "Get Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	Reactions, err := handler.service.Get(ctx, request.PostId)
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
	span := tracer.StartSpanFromContext(ctx, "Reaction Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	reaction := mapPbToReaction(request.Reaction)
	reaction, err := handler.service.Reaction(ctx, reaction)
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
	span := tracer.StartSpanFromContext(ctx, "Delete Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	err := handler.service.Delete(ctx, request.Id)
	if err != nil {
		log.Errorf("Cannot delete reaction: %v", err)
		return nil, err
	}
	log.Info("Reaction deleted")
	return &pb.DeleteResponse{}, nil
}

func (handler *ReactionHandler) DeletePostReactions(ctx context.Context,
	request *pb.DeletePostReactionsRequest) (*pb.DeletePostReactionsResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "DeletePostReactions Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	postId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		log.Errorf("Cannot parse postId: %v", err)
		return nil, err
	}
	err = handler.service.DeletePostReactions(ctx, postId)
	if err != nil {
		log.Errorf("Cannot delete reactions: %v", err)
		return nil, err
	}
	log.Info("Reactions deleted")
	return &pb.DeletePostReactionsResponse{}, nil
}

func (handler *ReactionHandler) GetLogs(ctx context.Context, request *pb.GetLogsRequest) (*pb.GetLogsResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "GetLogs Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	logs, err := handler.service.GetLogs(ctx)
	if err != nil {
		log.Errorf("GLF")
		return nil, err
	}
	pbLogs := make([]*pb.Log, len(logs))
	for i, log := range logs {
		pbLogs[i] = &pb.Log{
			Time:        timestamppb.New(log.Time),
			Level:       log.Level,
			Service:     "Reaction service",
			Msg:         log.Msg,
			FullContent: log.FullContent,
		}
	}
	log.Info("GLD")
	return &pb.GetLogsResponse{Logs: pbLogs}, nil
}
