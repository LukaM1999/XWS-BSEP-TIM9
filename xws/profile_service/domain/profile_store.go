package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProfileStore interface {
	Get(profileId string) (*Profile, error)
	GetAll(search string) ([]*Profile, error)
	Create(profile *Profile) error
	Update(profileId string, profile *Profile) error
	DeleteAll() error
	Delete(id string) error
	GetByToken(token string) (*Profile, error)
	GenerateToken(id primitive.ObjectID) (string, error)
}
