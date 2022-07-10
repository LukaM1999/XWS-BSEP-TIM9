package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostStore interface {
	Get(ctx context.Context, id string) (*Post, error)
	GetProfilePosts(ctx context.Context, profileId string) ([]*Post, error)
	GetConnectionPosts(ctx context.Context, profileId string) ([]*Post, error)
	Create(ctx context.Context, post *Post) error
	CreateConnection(ctx context.Context, connection *Connection) error
	DeleteConnection(ctx context.Context, id primitive.ObjectID) error
	Update(ctx context.Context, id string, post *Post) error
	UpdateProfile(ctx context.Context, id primitive.ObjectID, profile *Profile) error
	Delete(ctx context.Context, id string) error
	DeleteAll(ctx context.Context) error
	UpdatePostImage(ctx context.Context, id primitive.ObjectID, url string) (*Post, error)
}
