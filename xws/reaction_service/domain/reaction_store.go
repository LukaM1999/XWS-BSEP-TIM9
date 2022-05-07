package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReactionStore interface {
	Get(postId string) ([]*Reaction, error)
	Reaction(reaction *Reaction) (*Reaction, error)
	Delete(id string) error
	DeletePostReactions(postId primitive.ObjectID) error
	DeleteAll() error
}
