package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReactionStore interface {
	Get(ctx context.Context, postId string) ([]*Reaction, error)
	Reaction(ctx context.Context, reaction *Reaction) (*Reaction, error)
	Delete(ctx context.Context, id string) error
	DeletePostReactions(ctx context.Context, postId primitive.ObjectID) error
	DeleteAll(ctx context.Context) error
}
