package application

import (
	"context"
	auth "dislinkt/common/domain"
	"dislinkt/common/tracer"
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

func (service *ReactionService) Get(ctx context.Context, postId string) ([]*domain.Reaction, error) {
	span := tracer.StartSpanFromContext(ctx, "Get Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.Get(ctx, postId)
}

func (service *ReactionService) Reaction(ctx context.Context, reaction *domain.Reaction) (*domain.Reaction, error) {
	span := tracer.StartSpanFromContext(ctx, "Reaction Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.Reaction(ctx, reaction)
}

func (service *ReactionService) Delete(ctx context.Context, id string) error {
	span := tracer.StartSpanFromContext(ctx, "Delete Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.Delete(ctx, id)
}

func (service *ReactionService) DeletePostReactions(ctx context.Context, postId primitive.ObjectID) error {
	span := tracer.StartSpanFromContext(ctx, "DeletePostReactions Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.DeletePostReactions(ctx, postId)
}

func (service *ReactionService) GetLogs(ctx context.Context) ([]auth.Log, error) {
	span := tracer.StartSpanFromContext(ctx, "GetLogs Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	
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
