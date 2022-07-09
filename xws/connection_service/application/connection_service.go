package application

import (
	auth "dislinkt/common/domain"
	"dislinkt/connection_service/domain"
	"os"
	"regexp"
	"strings"
	"time"
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

func (service *ConnectionService) GetConnection(user1Id string, user2Id string) (*domain.Connection, error) {
	return service.store.GetConnection(user1Id, user2Id)
}

func (service *ConnectionService) GetLogs() ([]auth.Log, error) {
	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	content, err := os.ReadFile(logPathPrefix + "connection_service/connection.log")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	logs := make([]auth.Log, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		var log auth.Log
		splitBySpace := strings.Split(line, " ")
		log.Time, err = time.Parse("2006-01-02T15:04:05.000Z", strings.Trim(strings.Split(splitBySpace[0], "=")[1], "\""))
		if err != nil {
			log.Time = time.Time{}
		}
		log.Level = strings.Split(splitBySpace[1], "=")[1]
		re := regexp.MustCompile(`msg="[/\\=!?'"\.a-zA-Z0-9_\s:-]*"`)
		msg := re.FindString(line)
		if msg != "" {
			log.Msg = strings.Trim(strings.Split(msg, "=")[1], "\"")
		}
		if msg == "" {
			re = regexp.MustCompile(`msg=[a-zA-Z]*`)
			msg = re.FindString(line)
			if msg != "" {
				log.Msg = strings.Split(msg, "=")[1]
			}
		}
		log.FullContent = line
		logs = append(logs, log)
	}
	return logs, nil
}
