package application

import (
	"context"
	auth "dislinkt/common/domain"
	"dislinkt/common/tracer"
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

func (service *ConnectionService) Get(ctx context.Context, userId string) ([]*domain.Connection, error) {
	span := tracer.StartSpanFromContext(ctx, "Get Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.Get(ctx, userId)
}

func (service *ConnectionService) Create(ctx context.Context, issuerKey string, subjectKey string) (*domain.Connection, error) {
	span := tracer.StartSpanFromContext(ctx, "Create Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.CreateConnection(ctx, issuerKey, subjectKey)
}

func (service *ConnectionService) CreateUser(ctx context.Context, userId string) error {
	span := tracer.StartSpanFromContext(ctx, "CreateUser Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.CreateUser(ctx, userId)
}

func (service *ConnectionService) Delete(ctx context.Context, id int) error {
	span := tracer.StartSpanFromContext(ctx, "Delete Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.Delete(ctx, id)
}

func (service *ConnectionService) DeleteUser(ctx context.Context, id string) error {
	span := tracer.StartSpanFromContext(ctx, "DeleteUser Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.DeleteUser(ctx, id)
}

func (service *ConnectionService) UpdateConnection(ctx context.Context, id int) (*domain.Connection, error) {
	span := tracer.StartSpanFromContext(ctx, "UpdateConnection Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.UpdateConnection(ctx, id)
}

func (service *ConnectionService) UpdatePrivacy(ctx context.Context, id string) error {
	span := tracer.StartSpanFromContext(ctx, "UpdatePrivacy Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.UpdatePrivacy(ctx, id)
}

func (service *ConnectionService) GetRecommendations(ctx context.Context, userId string) ([]string, error) {
	span := tracer.StartSpanFromContext(ctx, "GetRecommendations Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.GetRecommendations(ctx, userId)
}

func (service *ConnectionService) BlockUser(ctx context.Context, issuerId string, subjectId string) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "BlockUser Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.BlockUser(ctx, issuerId, subjectId)
}

func (service *ConnectionService) GetBlockedUsers(ctx context.Context, userId string) ([]string, error) {
	span := tracer.StartSpanFromContext(ctx, "GetBlockedUsers Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.GetBlockedUsers(ctx, userId)
}

func (service *ConnectionService) GetBlockers(ctx context.Context, userId string) ([]string, error) {
	span := tracer.StartSpanFromContext(ctx, "GetBlockers Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.GetBlockers(ctx, userId)
}

func (service *ConnectionService) UnblockUser(ctx context.Context, issuerId, subjectId string) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "UnblockUser Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.UnblockUser(ctx, issuerId, subjectId)
}

func (service *ConnectionService) GetConnection(ctx context.Context, user1Id string, user2Id string) (*domain.Connection, error) {
	span := tracer.StartSpanFromContext(ctx, "GetConnection Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.GetConnection(ctx, user1Id, user2Id)
}

func (service *ConnectionService) GetLogs(ctx context.Context) ([]auth.Log, error) {
	span := tracer.StartSpanFromContext(ctx, "GetLogs Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	
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
