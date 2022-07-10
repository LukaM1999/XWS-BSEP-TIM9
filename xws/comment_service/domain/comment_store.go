package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentStore interface {
	Get(ctx context.Context, postId string) ([]*Comment, error)
	Create(ctx context.Context, comment *Comment) (*Comment, error)
	Delete(ctx context.Context, id string) error
	UpdateCommentCreator(ctx context.Context, creatorId primitive.ObjectID, creator *CommentCreator) error
	DeletePostComments(ctx context.Context, postId primitive.ObjectID) error
	DeleteAll(ctx context.Context) error
}
