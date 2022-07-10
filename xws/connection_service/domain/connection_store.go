package domain

import "context"

type ConnectionStore interface {
	Get(ctx context.Context, userId string) ([]*Connection, error)
	CreateUser(ctx context.Context, primaryKey string) error
	CreateConnection(ctx context.Context, issuerKey string, subjectKey string) (*Connection, error)
	Delete(ctx context.Context, id int) error
	DeleteUser(ctx context.Context, userId string) error
	DeleteAll(ctx context.Context) error
	UpdateConnection(ctx context.Context, id int) (*Connection, error)
	UpdatePrivacy(ctx context.Context, userKey string) error
	GetRecommendations(ctx context.Context, userId string) ([]string, error)
	BlockUser(ctx context.Context, issuerId string, subjectId string) (bool, error)
	GetBlockedUsers(ctx context.Context, userId string) ([]string, error)
	GetBlockers(ctx context.Context, userId string) ([]string, error)
	UnblockUser(ctx context.Context, issuerId string, subjectId string) (bool, error)
	GetConnection(ctx context.Context, user1Id string, user2Id string) (*Connection, error)
}
