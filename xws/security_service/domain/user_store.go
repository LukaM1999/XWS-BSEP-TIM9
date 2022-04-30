package domain

type UserStore interface {
	Get(username string) (*User, error)
	GetAll() ([]*User, error)
	Register(user *User) error
	DeleteAll() error
}
