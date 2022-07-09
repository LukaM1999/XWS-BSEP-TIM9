package application

import (
	"dislinkt/comment_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"strings"
)

type CommentService struct {
	store domain.CommentStore
}

func NewCommentService(store domain.CommentStore) *CommentService {
	return &CommentService{
		store: store,
	}
}

func (service *CommentService) Get(postId string) ([]*domain.Comment, error) {
	return service.store.Get(postId)
}

func (service *CommentService) Create(comment *domain.Comment) (*domain.Comment, error) {
	return service.store.Create(comment)
}

func (service *CommentService) Delete(id string) error {
	return service.store.Delete(id)
}

func (service *CommentService) UpdateCommentCreator(creatorId primitive.ObjectID, creator *domain.CommentCreator) error {
	return service.store.UpdateCommentCreator(creatorId, creator)
}

func (service *CommentService) DeletePostComments(postId primitive.ObjectID) error {
	return service.store.DeletePostComments(postId)
}

func (service *CommentService) GetLogs() ([]string, error) {
	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	content, err := os.ReadFile(logPathPrefix + "comment_service/comment.log")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}
