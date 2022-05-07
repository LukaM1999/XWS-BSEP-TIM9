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

func (service *ProfileService) Get(profileId string) (*domain.Profile, error) {
	return service.store.Get(profileId)
}

func (service *ProfileService) GetAll(search string) ([]*domain.Profile, error) {
	return service.store.GetAll(search)
}

func (service *ProfileService) Create(profile *domain.Profile) error {
	return service.store.Create(profile)
}

func (service *ProfileService) Update(profileId string, profile *domain.Profile) error {
	return service.store.Update(profileId, profile)
}
