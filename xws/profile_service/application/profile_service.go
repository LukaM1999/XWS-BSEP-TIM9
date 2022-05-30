package application

import (
	"dislinkt/profile_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (service *ProfileService) Delete(id string) error {
	return service.store.Delete(id)
}

func (service *ProfileService) GetByToken(token string) (*domain.Profile, error) {
	return service.store.GetByToken(token)
}

func (service *ProfileService) GenerateToken(profileId string) (string, error) {
	id, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return "", err
	}

	return service.store.GenerateToken(id)
}
