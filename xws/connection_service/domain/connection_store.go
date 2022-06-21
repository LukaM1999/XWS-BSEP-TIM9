package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ConnectionStore interface {
	Get(userId string) ([]*Connection, error)
	Create(connection *Connection) (*Connection, error)
	CreatePrivacy(privacy *ProfilePrivacy) (*ProfilePrivacy, error)
	Delete(id string) error
	DeleteAll() error
	Update(id string) (*Connection, error)
	CreateProfilePrivacy(privacy *ProfilePrivacy) (*ProfilePrivacy, error)
	DeleteProfilePrivacy(id primitive.ObjectID) error
	UpdatePrivacy(id primitive.ObjectID) error
}
