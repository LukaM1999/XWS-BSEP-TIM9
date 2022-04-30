package application

import (
	"dislinkt/security_service/domain"
)

type SecurityService struct {
	store domain.UserStore
}

func NewSecurityService(store domain.UserStore) *SecurityService {
	return &SecurityService{
		store: store,
	}
}

func (service *SecurityService) Get(username string) (*domain.User, error) {
	return service.store.Get(username)
}

func (service *SecurityService) GetAll() ([]*domain.User, error) {
	return service.store.GetAll()
}

func (service *SecurityService) Register(user *domain.User) error {
	return service.store.Register(user)
}
