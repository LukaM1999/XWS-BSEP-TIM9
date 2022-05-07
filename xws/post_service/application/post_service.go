package application

import (
	"dislinkt/post_service/domain"
)

type PostService struct {
	store domain.PostStore
}

func NewPostService(store domain.PostStore) *PostService {
	return &PostService{
		store: store,
	}
}

func (service *PostService) Get(id string) (*domain.Post, error) {
	return service.store.Get(id)
}

func (service *PostService) GetProfilePosts(profileId string) ([]*domain.Post, error) {
	return service.store.GetProfilePosts(profileId)
}

func (service *PostService) GetConnectionPosts(profileId string) ([]*domain.Post, error) {
	return service.store.GetConnectionPosts(profileId)
}

func (service *PostService) Create(profile *domain.Post) error {
	return service.store.Create(profile)
}

func (service *PostService) Update(id string, post *domain.Post) error {
	return service.store.Update(id, post)
}

func (service *PostService) Delete(id string) error {
	return service.store.Delete(id)
}
