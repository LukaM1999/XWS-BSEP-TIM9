package application

import (
	"context"
	auth "dislinkt/common/domain"
	pbProfile "dislinkt/common/proto/profile_service"
	"dislinkt/common/tracer"
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

func (service *PostService) Get(ctx context.Context, id string) (*domain.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "Get Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.Get(ctx, id)
}

func (service *PostService) GetProfilePosts(ctx context.Context, profileId string) ([]*domain.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "GetProfilePosts Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.GetProfilePosts(ctx, profileId)
}

func (service *PostService) GetConnectionPosts(ctx context.Context, profileId string) ([]*domain.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "GetConnectionPosts Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.GetConnectionPosts(ctx, profileId)
}

func (service *PostService) Create(ctx context.Context, profile *domain.Post) error {
	span := tracer.StartSpanFromContext(ctx, "Create Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.Create(ctx, profile)
}

func (service *PostService) Update(ctx context.Context, id string, post *domain.Post) error {
	span := tracer.StartSpanFromContext(ctx, "Update Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.Update(ctx, id, post)
}

func (service *PostService) UpdateProfile(ctx context.Context, id primitive.ObjectID, profile *domain.Profile) error {
	span := tracer.StartSpanFromContext(ctx, "UpdateProfile Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.UpdateProfile(ctx, id, profile)
}

func (service *PostService) Delete(ctx context.Context, id string) error {
	span := tracer.StartSpanFromContext(ctx, "Delete Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.store.Delete(ctx, id)
	if err != nil {
		return err
	}
	err = service.orchestrator.Start(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *PostService) CreateConnection(ctx context.Context, connection *domain.Connection) error {
	span := tracer.StartSpanFromContext(ctx, "CreateConnection Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.CreateConnection(ctx, connection)
}

func (service *PostService) DeleteConnection(ctx context.Context, id primitive.ObjectID) error {
	span := tracer.StartSpanFromContext(ctx, "DeleteConnection Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.DeleteConnection(ctx, id)
}

func (service *PostService) GetLogs(ctx context.Context) ([]auth.Log, error) {
	span := tracer.StartSpanFromContext(ctx, "GetLogs Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

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

func (service *PostService) UpdatePostImage(ctx context.Context, id primitive.ObjectID, url string) (*domain.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "UpdatePostImage Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.UpdatePostImage(ctx, id, url)
}
