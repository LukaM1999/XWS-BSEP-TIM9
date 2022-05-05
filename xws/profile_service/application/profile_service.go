package application

import (
	"dislinkt/profile_service/domain"
)

type ProfileService struct {
	store domain.ProfileStore
}

func NewProfileService(store domain.ProfileStore) *ProfileService {
	return &ProfileService{
		store: store,
	}
}

func (service *ProfileService) Get(username string) (*domain.Profile, error) {
	return service.store.Get(username)
}

func (service *ProfileService) GetAll(search string) ([]*domain.Profile, error) {
	return service.store.GetAll(search)
}

func (service *ProfileService) Create(profile *domain.Profile) error {
	return service.store.Create(profile)
}

func (service *ProfileService) Update(username string, profile *domain.Profile) error {
	return service.store.Update(username, profile)
}
