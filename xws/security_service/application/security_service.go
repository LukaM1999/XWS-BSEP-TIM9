package application

import (
	auth "dislinkt/common/domain"
	"dislinkt/security_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SecurityService struct {
	store domain.UserStore
}

func NewSecurityService(store domain.UserStore) *SecurityService {
	return &SecurityService{
		store: store,
	}
}

func (service *SecurityService) Get(username string) (*auth.User, error) {
	return service.store.Get(username)
}

func (service *SecurityService) GetAll() ([]*auth.User, error) {
	return service.store.GetAll()
}

func (service *SecurityService) Register(user *auth.User) (*auth.User, error) {
	return service.store.Register(user)
}

func (service *SecurityService) Update(id primitive.ObjectID, username string) (string, error) {
	return service.store.Update(id, username)
}

func (service *SecurityService) Delete(id primitive.ObjectID) error {
	return service.store.Delete(id)
}
