package application

import (
	"context"
	"dislinkt/comment_service/domain"
	auth "dislinkt/common/domain"
	"dislinkt/common/tracer"
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

func (service *CommentService) Get(ctx context.Context, postId string) ([]*domain.Comment, error) {
	span := tracer.StartSpanFromContext(ctx, "Get Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.Get(ctx, postId)
}

func (service *CommentService) Create(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	span := tracer.StartSpanFromContext(ctx, "Create Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.Create(ctx, comment)
}

func (service *CommentService) Delete(ctx context.Context, id string) error {
	span := tracer.StartSpanFromContext(ctx, "Delete Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.Delete(ctx, id)
}

func (service *CommentService) UpdateCommentCreator(ctx context.Context, creatorId primitive.ObjectID, creator *domain.CommentCreator) error {
	span := tracer.StartSpanFromContext(ctx, "UpdateCommentCreator Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.UpdateCommentCreator(ctx, creatorId, creator)
}

func (service *CommentService) DeletePostComments(ctx context.Context, postId primitive.ObjectID) error {
	span := tracer.StartSpanFromContext(ctx, "DeletePostComments Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.DeletePostComments(ctx, postId)
}

func (service *CommentService) GetLogs(ctx context.Context, ) ([]auth.Log, error) {
	span := tracer.StartSpanFromContext(ctx, "GetLogs Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

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
