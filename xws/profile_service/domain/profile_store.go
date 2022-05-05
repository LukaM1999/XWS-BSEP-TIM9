package domain

type ProfileStore interface {
	Get(username string) (*Profile, error)
	GetAll(search string) ([]*Profile, error)
	Create(profile *Profile) error
	Update(username string, profile *Profile) error
	DeleteAll() error
}
