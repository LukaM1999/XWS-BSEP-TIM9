package application

import (
	"dislinkt/reaction_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"strings"
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

func (service *ReactionService) GetLogs() ([]string, error) {
	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	content, err := os.ReadFile(logPathPrefix + "reaction_service/reaction.log")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}
