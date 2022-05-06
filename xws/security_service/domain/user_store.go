package domain

import auth "dislinkt/common/domain"

type UserStore interface {
	Get(username string) (*auth.User, error)
	GetAll() ([]*auth.User, error)
	Register(user *auth.User) error
	DeleteAll() error
}
