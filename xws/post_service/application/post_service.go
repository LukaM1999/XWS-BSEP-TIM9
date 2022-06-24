package application

import (
	"context"
	pbProfile "dislinkt/common/proto/profile_service"
	"dislinkt/post_service/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostService struct {
	store         domain.PostStore
	profileClient pbProfile.ProfileServiceClient
	orchestrator  *DeletePostOrchestrator
}

func NewPostService(store domain.PostStore, profileClient pbProfile.ProfileServiceClient, orchestrator *DeletePostOrchestrator) *PostService {
	return &PostService{
		store:         store,
		profileClient: profileClient,
		orchestrator:  orchestrator,
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
	err := service.store.Delete(id)
	if err != nil {
		return err
	}
	err = service.orchestrator.Start(id)
	if err != nil {
		return err
	}
	return nil
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

func (service *PostService) PromoteJob(job *domain.JobOffer, token string, username string) (*domain.JobOffer, error) {
	profile, err := service.profileClient.GetByToken(context.TODO(), &pbProfile.GetByTokenRequest{Token: &token})
	if err != nil {
		return nil, err
	}
	if *profile.Profile.Username != username {
		return nil, errors.New("invalid username or token")
	}
	return service.store.CreateJob(job)
}
