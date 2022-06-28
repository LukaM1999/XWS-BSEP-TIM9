package domain

type ConnectionStore interface {
	Get(userId string) ([]*Connection, error)
	CreateUser(primaryKey string) error
	CreateConnection(issuerKey string, subjectKey string) (*Connection, error)
	Delete(id int) error
	DeleteUser(userId string) error
	DeleteAll() error
	UpdateConnection(id int) (*Connection, error)
	UpdatePrivacy(userKey string) error
	GetRecommendations(userId string) ([]string, error)
	BlockUser(issuerId string, subjectId string) (bool, error)
	GetBlockedUsers(userId string) ([]string, error)
	GetBlockers(userId string) ([]string, error)
	UnblockUser(issuerId string, subjectId string) (bool, error)
}
