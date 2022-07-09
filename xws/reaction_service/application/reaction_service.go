package application

import (
	auth "dislinkt/common/domain"
	"dislinkt/reaction_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"regexp"
	"strings"
	"time"
)

type ReactionService struct {
	store domain.ReactionStore
}

func NewReactionService(store domain.ReactionStore) *ReactionService {
	return &ReactionService{
		store: store,
	}
}

func (service *ReactionService) Get(postId string) ([]*domain.Reaction, error) {
	return service.store.Get(postId)
}

func (service *ReactionService) Reaction(reaction *domain.Reaction) (*domain.Reaction, error) {
	return service.store.Reaction(reaction)
}

func (service *ReactionService) Delete(id string) error {
	return service.store.Delete(id)
}

func (service *ReactionService) DeletePostReactions(postId primitive.ObjectID) error {
	return service.store.DeletePostReactions(postId)
}

func (service *ReactionService) GetLogs() ([]auth.Log, error) {
	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	content, err := os.ReadFile(logPathPrefix + "reaction_service/reaction.log")
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
