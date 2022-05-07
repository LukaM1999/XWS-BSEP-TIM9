package api

import (
	"dislinkt/comment_service/domain"
	pb "dislinkt/common/proto/comment_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapCommentToPb(comment *domain.Comment) *pb.Comment {
	pbComment := &pb.Comment{
		Id:      comment.Id.Hex(),
		Content: comment.Content,
		CommentCreator: &pb.CommentCreator{
			Id:        comment.CommentCreator.Id.Hex(),
			FirstName: comment.CommentCreator.FirstName,
			LastName:  comment.CommentCreator.LastName,
		},
		PostId:      comment.PostId.Hex(),
		DateCreated: timestamppb.New(comment.DateCreated),
	}

	return pbComment
}

func mapPbToComment(pbComment *pb.Comment) *domain.Comment {
	comment := &domain.Comment{
		Id:      getObjectId(pbComment.Id),
		Content: pbComment.Content,
		CommentCreator: domain.CommentCreator{
			Id:        getObjectId(pbComment.CommentCreator.Id),
			FirstName: pbComment.CommentCreator.FirstName,
			LastName:  pbComment.CommentCreator.LastName,
		},
		PostId:      getObjectId(pbComment.PostId),
		DateCreated: pbComment.DateCreated.AsTime(),
	}

	return comment
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
