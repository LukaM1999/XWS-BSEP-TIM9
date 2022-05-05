package application

import (
	"dislinkt/connection_service/domain"
)

type ConnectionService struct {
	store domain.ConnectionStore
}

func NewConnectionService(store domain.ConnectionStore) *ConnectionService {
	return &ConnectionService{
		store: store,
	}
}

func (service *ConnectionService) Get(userId string) ([]*domain.Connection, error) {
	return service.store.Get(userId)
}

func (service *ConnectionService) Create(connection *domain.Connection) (*domain.Connection, error) {
	return service.store.Create(connection)
}

func (service *ConnectionService) Delete(id string) error {
	return service.store.Delete(id)
}

func (service *ConnectionService) Update(id string) (*domain.Connection, error) {
	return service.store.Update(id)
}
