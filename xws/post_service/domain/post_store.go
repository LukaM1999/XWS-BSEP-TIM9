package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type PostStore interface {
	Get(id string) (*Post, error)
	GetProfilePosts(profileId string) ([]*Post, error)
	GetConnectionPosts(profileId string) ([]*Post, error)
	Create(post *Post) error
	CreateConnection(connection *Connection) error
	DeleteConnection(id primitive.ObjectID) error
	Update(id string, post *Post) error
	UpdateProfile(id primitive.ObjectID, profile *Profile) error
	Delete(id string) error
	DeleteAll() error
	UpdatePostImage(id primitive.ObjectID, url string) (*Post, error)
}
