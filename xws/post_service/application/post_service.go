package application

import (
	auth "dislinkt/common/domain"
	pbProfile "dislinkt/common/proto/profile_service"
	"dislinkt/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"regexp"
	"strings"
	"time"
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

func (service *PostService) GetLogs() ([]auth.Log, error) {
	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	content, err := os.ReadFile(logPathPrefix + "post_service/post.log")
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
