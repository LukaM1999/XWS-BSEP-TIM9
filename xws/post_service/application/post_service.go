package application

import (
	"dislinkt/post_service/domain"
	app "dislinkt/profile_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostService struct {
	store          domain.PostStore
	profileService app.ProfileService
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

func (service *PostService) UpdateProfile(id primitive.ObjectID, profile *domain.Profile) error {
	return service.store.UpdateProfile(id, profile)
}

func (service *PostService) Delete(id string) error {
	return service.store.Delete(id)
}

func (service *PostService) CreateConnection(connection *domain.Connection) error {
	return service.store.CreateConnection(connection)
}

func (service *PostService) DeleteConnection(id primitive.ObjectID) error {
	return service.store.DeleteConnection(id)
}

func (service *PostService) CreateJob(job *domain.JobOffer) (*domain.JobOffer, error) {
	return service.store.CreateJob(job)
}

func (service *PostService) PromoteJob(job *domain.JobOffer, token string) (*domain.JobOffer, error) {
	_, err := service.profileService.GetByToken(token)
	if err != nil {
		return nil, err
	}
	return service.store.CreateJob(job)
}
