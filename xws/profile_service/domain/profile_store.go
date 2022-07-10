package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileStore interface {
	Get(ctx context.Context, profileId string) (*Profile, error)
	GetAll(ctx context.Context, search string) ([]*Profile, error)
	Create(ctx context.Context, profile *Profile) error
	Update(ctx context.Context, profileId string, profile *Profile) error
	DeleteAll(ctx context.Context) error
	Delete(ctx context.Context, id string) error
	GetByToken(ctx context.Context, token string) (*Profile, error)
	GenerateToken(ctx context.Context, id primitive.ObjectID) (string, error)
}
