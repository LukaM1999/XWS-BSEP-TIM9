package domain

import (
	auth "dislinkt/common/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStore interface {
	Get(username string) (*auth.User, error)
	GetAll() ([]*auth.User, error)
	Register(user *auth.User) (*auth.User, error)
	Update(id primitive.ObjectID, username string) (string, error)
	Delete(id primitive.ObjectID) error
	DeleteAll() error
}
