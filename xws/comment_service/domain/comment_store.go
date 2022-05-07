package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type CommentStore interface {
	Get(postId string) ([]*Comment, error)
	Create(comment *Comment) (*Comment, error)
	Delete(id string) error
	UpdateCommentCreator(creatorId primitive.ObjectID, creator *CommentCreator) error
	DeletePostComments(postId primitive.ObjectID) error
	DeleteAll() error
}
