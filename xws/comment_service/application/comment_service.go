package application

import (
	"dislinkt/comment_service/domain"
	auth "dislinkt/common/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"regexp"
	"strings"
	"time"
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

func (service *CommentService) GetLogs() ([]auth.Log, error) {
	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	content, err := os.ReadFile(logPathPrefix + "comment_service/comment.log")
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
