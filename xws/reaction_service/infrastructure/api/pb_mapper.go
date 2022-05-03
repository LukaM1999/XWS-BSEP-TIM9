package api

import (
	pb "dislinkt/common/proto/reaction_service"
	"dislinkt/reaction_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapReactionToPb(reaction *domain.Reaction) *pb.Reaction {
	return &pb.Reaction{
		Id:        reaction.Id.Hex(),
		UserId:    reaction.UserId.Hex(),
		Type:      pb.Reaction_ReactionType(reaction.Type),
		PostId:    reaction.PostId.Hex(),
		CreatedAt: timestamppb.New(reaction.CreatedAt),
	}
}

func mapPbToReaction(pbReaction *pb.Reaction) *domain.Reaction {
	return &domain.Reaction{
		Id:        getObjectId(pbReaction.Id),
		UserId:    getObjectId(pbReaction.UserId),
		Type:      domain.ReactionType(pbReaction.Type),
		PostId:    getObjectId(pbReaction.PostId),
		CreatedAt: pbReaction.CreatedAt.AsTime(),
	}
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
