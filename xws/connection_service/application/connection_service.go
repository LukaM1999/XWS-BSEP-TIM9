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

func (service *ConnectionService) Create(issuerKey string, subjectKey string) (*domain.Connection, error) {
	return service.store.CreateConnection(issuerKey, subjectKey)
}

func (service *ConnectionService) CreateUser(userId string) error {
	return service.store.CreateUser(userId)
}

func (service *ConnectionService) Delete(id int) error {
	return service.store.Delete(id)
}

func (service *ConnectionService) DeleteUser(id string) error {
	return service.store.DeleteUser(id)
}

func (service *ConnectionService) UpdateConnection(id int) (*domain.Connection, error) {
	return service.store.UpdateConnection(id)
}

func (service *ConnectionService) UpdatePrivacy(id string) error {
	return service.store.UpdatePrivacy(id)
}

func (service *ConnectionService) GetRecommendations(userId string) ([]string, error) {
	return service.store.GetRecommendations(userId)
}

func (service *ConnectionService) BlockUser(issuerId string, subjectId string) (bool, error) {
	return service.store.BlockUser(issuerId, subjectId)
}

func (service *ConnectionService) GetBlockedUsers(userId string) ([]string, error) {
	return service.store.GetBlockedUsers(userId)
}

func (service *ConnectionService) GetBlockers(userId string) ([]string, error) {
	return service.store.GetBlockers(userId)
}

func (service *ConnectionService) UnblockUser(issuerId, subjectId string) (bool, error) {
	return service.store.UnblockUser(issuerId, subjectId)
}
