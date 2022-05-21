package domain

type ProfileStore interface {
	Get(profileId string) (*Profile, error)
	GetAll(search string) ([]*Profile, error)
	Create(profile *Profile) error
	Update(profileId string, profile *Profile) error
	DeleteAll() error
	Delete(id string) error
}
