package persistence

import (
	"context"
	"dislinkt/connection_service/domain"
	"dislinkt/connection_service/ent"
	"dislinkt/connection_service/ent/blockeduser"
	"dislinkt/connection_service/ent/connection"
	"dislinkt/connection_service/ent/user"
	"fmt"
	"log"
)

type ConnectionPostgresStore struct {
	connectionString string
}

func NewConnectionPostgresStore(host string, port string) domain.ConnectionStore {
	connectionString := fmt.Sprintf("host=%s port=%s user=postgres password=ftn dbname=dislinkt sslmode=disable", host, port)
	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Schema.Create(context.TODO()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	//err = entc.Generate("./ent/schema", &gen.Config{}, entc.Extensions(entviz.Extension{}))
	//if err != nil {
	//	log.Fatalf("running ent codegen: %v", err)
	//}
	return &ConnectionPostgresStore{
		connectionString: connectionString,
	}
}

func (store *ConnectionPostgresStore) Get(userId string) ([]*domain.Connection, error) {
	client, err := ent.Open("postgres", store.connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	all, err := client.Connection.Query().
		Where(connection.And(connection.HasUserWith(user.PrimaryKeyEQ(userId))),
			connection.IsApprovedEQ(true)).
		All(context.TODO())
	if err != nil {
		return nil, err
	}
	var connections = make([]*domain.Connection, 0)
	for _, connection := range all {
		connections = append(connections, &domain.Connection{
			Id:         connection.ID,
			IssuerId:   connection.IssuerPrimaryKey,
			SubjectId:  connection.SubjectPrimaryKey,
			Date:       connection.CreatedAt,
			IsApproved: connection.IsApproved,
		})
	}
	return connections, nil
}

func (store *ConnectionPostgresStore) CreateUser(primaryKey string) error {
	client, err := ent.Open("postgres", store.connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	_, err = client.User.
		Create().
		SetPrimaryKey(primaryKey).
		Save(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectionPostgresStore) CreateConnection(issuerKey string, subjectKey string) (*domain.Connection, error) {
	client, err := ent.Open("postgres", store.connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)
	issuer, err := client.User.Query().Where(user.PrimaryKeyEQ(issuerKey)).Only(context.TODO())
	if err != nil {
		return nil, err
	}
	subject, err := client.User.Query().Where(user.PrimaryKeyEQ(subjectKey)).Only(context.TODO())
	if err != nil {
		return nil, err
	}
	newConnection, err := client.Connection.
		Create().
		SetIssuerPrimaryKey(issuerKey).
		SetSubjectPrimaryKey(subjectKey).
		SetIsApproved(!subject.IsPrivate).
		SetUserID(issuer.ID).
		SetConnectionID(subject.ID).
		Save(context.TODO())

	if err != nil {
		return nil, err
	}
	return &domain.Connection{
		Id:         newConnection.ID,
		IssuerId:   newConnection.IssuerPrimaryKey,
		SubjectId:  newConnection.SubjectPrimaryKey,
		Date:       newConnection.CreatedAt,
		IsApproved: newConnection.IsApproved,
	}, nil

}

func (store *ConnectionPostgresStore) DeleteAll() error {
	client, err := ent.Open("postgres", store.connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	_, err = client.Connection.Delete().Exec(context.TODO())
	if err != nil {
		return err
	}
	_, err = client.BlockedUser.Delete().Exec(context.TODO())
	if err != nil {
		return err
	}
	_, err = client.User.Delete().Exec(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectionPostgresStore) Delete(id int) error {
	client, err := ent.Open("postgres", store.connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	err = client.Connection.DeleteOneID(id).Exec(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectionPostgresStore) DeleteUser(userId string) error {
	client, err := ent.Open("postgres", store.connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	_, err = client.User.Delete().Where(user.PrimaryKeyEQ(userId)).Exec(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectionPostgresStore) UpdateConnection(id int) (*domain.Connection, error) {
	client, err := ent.Open("postgres", store.connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	oldConnection, err := client.Connection.Query().Where(connection.ConnectionID(id)).Only(context.TODO())
	if err != nil {
		return nil, err
	}
	newConnection, err := oldConnection.Update().SetIsApproved(!oldConnection.IsApproved).Save(context.TODO())
	if err != nil {
		return nil, err
	}
	return &domain.Connection{
		Id:         newConnection.ID,
		IssuerId:   newConnection.IssuerPrimaryKey,
		SubjectId:  newConnection.SubjectPrimaryKey,
		Date:       newConnection.CreatedAt,
		IsApproved: newConnection.IsApproved,
	}, nil
}

func (store *ConnectionPostgresStore) UpdatePrivacy(userKey string) error {
	client, err := ent.Open("postgres", store.connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	oldUser, err := client.User.Query().Where(user.PrimaryKeyEQ(userKey)).Only(context.TODO())
	if err != nil {
		return err
	}
	newUser, err := oldUser.Update().SetIsPrivate(!oldUser.IsPrivate).Save(context.TODO())
	if err != nil {
		return err
	}
	if newUser.IsPrivate {
		return nil
	}
	client.Connection.Update().
		Where(connection.SubjectPrimaryKeyEQ(userKey), connection.IsApprovedEQ(false)).
		SetIsApproved(true).
		Save(context.TODO())
	return nil
}

func (store *ConnectionPostgresStore) GetRecommendations(userId string) ([]string, error) {
	client, err := ent.Open("postgres", store.connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	userConnectionIssuers, err := client.Connection.Query().
		Where(connection.And(connection.Or(connection.HasUserWith(user.PrimaryKeyEQ(userId)),
			connection.HasConnectionWith(user.PrimaryKeyEQ(userId))),
			connection.IsApprovedEQ(true))).
		QueryUser().Where(user.PrimaryKeyNEQ(userId)).
		All(context.TODO())
	if err != nil {
		return nil, err
	}
	userConnectionSubjects, err := client.Connection.Query().
		Where(connection.And(connection.Or(connection.HasUserWith(user.PrimaryKeyEQ(userId)),
			connection.HasConnectionWith(user.PrimaryKeyEQ(userId))),
			connection.IsApprovedEQ(true))).
		QueryConnection().Where(user.PrimaryKeyNEQ(userId)).
		All(context.TODO())
	if err != nil {
		return nil, err
	}
	var userConnections []*ent.User
	userConnections = append(userConnectionIssuers, userConnectionSubjects...)
	var recommendations = make(map[string]bool)
	for _, connectedUser := range userConnections {
		recommendationIssuers, err := client.Connection.Query().
			Where(connection.Or(connection.And(connection.HasUserWith(user.PrimaryKeyEQ(connectedUser.PrimaryKey)),
				connection.Not(connection.HasConnectionWith(user.PrimaryKeyEQ(userId)))),
				connection.And(connection.HasConnectionWith(user.PrimaryKeyEQ(connectedUser.PrimaryKey)),
					connection.Not(connection.HasUserWith(user.PrimaryKeyEQ(userId))))),
				connection.IsApprovedEQ(true)).
			QueryUser().Where(user.PrimaryKeyNEQ(connectedUser.PrimaryKey)).
			Select(user.FieldPrimaryKey).
			Strings(context.TODO())
		if err != nil {
			return nil, err
		}
		recommendationSubjects, err := client.Connection.Query().
			Where(connection.Or(connection.And(connection.HasUserWith(user.PrimaryKeyEQ(connectedUser.PrimaryKey)),
				connection.Not(connection.HasConnectionWith(user.PrimaryKeyEQ(userId)))),
				connection.And(connection.HasConnectionWith(user.PrimaryKeyEQ(connectedUser.PrimaryKey)),
					connection.Not(connection.HasUserWith(user.PrimaryKeyEQ(userId))))),
				connection.IsApprovedEQ(true)).
			QueryConnection().Where(user.PrimaryKeyNEQ(connectedUser.PrimaryKey)).
			Select(user.FieldPrimaryKey).
			Strings(context.TODO())
		if err != nil {
			return nil, err
		}
		for _, recommendationId := range recommendationIssuers {
			recommendations[recommendationId] = true
		}
		for _, recommendationId := range recommendationSubjects {
			recommendations[recommendationId] = true
		}

	}
	if err != nil {
		return nil, err
	}
	var recommendationIds = make([]string, 0, len(recommendations))
	for recommendationId := range recommendations {
		recommendationIds = append(recommendationIds, recommendationId)
	}
	return recommendationIds, nil
}

func (store *ConnectionPostgresStore) BlockUser(issuerId string, subjectId string) (bool, error) {
	client, err := ent.Open("postgres", store.connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	issuer, err := client.User.Query().Where(user.PrimaryKeyEQ(issuerId)).Only(context.TODO())
	if err != nil {
		return false, err
	}
	_, err = client.BlockedUser.Create().
		SetBlockedByID(issuer.ID).
		SetIssuerPrimaryKey(issuer.PrimaryKey).
		SetSubjectPrimaryKey(subjectId).
		Save(context.TODO())
	if err != nil {
		return false, err
	}
	_, err = client.Connection.Delete().
		Where(connection.Or(connection.And(connection.HasUserWith(user.PrimaryKeyEQ(issuerId)),
			connection.HasConnectionWith(user.PrimaryKeyEQ(subjectId))),
			connection.And(connection.HasUserWith(user.PrimaryKeyEQ(subjectId)),
				connection.HasConnectionWith(user.PrimaryKeyEQ(issuerId)))),
			connection.IsApprovedEQ(true)).
		Exec(context.TODO())
	if err != nil {
		return false, err
	}
	return true, nil
}

func (store *ConnectionPostgresStore) GetBlockedUsers(userId string) ([]string, error) {
	client, err := ent.Open("postgres", store.connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	blockedUsers, err := client.BlockedUser.Query().
		Where(blockeduser.IssuerPrimaryKey(userId)).
		Select(blockeduser.FieldSubjectPrimaryKey).
		Strings(context.TODO())
	if err != nil {
		return nil, err
	}
	return blockedUsers, nil
}

func (store *ConnectionPostgresStore) GetBlockers(userId string) ([]string, error) {
	client, err := ent.Open("postgres", store.connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	blockers, err := client.BlockedUser.Query().
		Where(blockeduser.SubjectPrimaryKey(userId)).
		Select(blockeduser.FieldIssuerPrimaryKey).
		Strings(context.TODO())
	if err != nil {
		return nil, err
	}
	return blockers, nil
}

func (store *ConnectionPostgresStore) UnblockUser(issuerId string, subjectId string) (bool, error) {
	client, err := ent.Open("postgres", store.connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	_, err = client.BlockedUser.Delete().
		Where(blockeduser.And(blockeduser.IssuerPrimaryKey(issuerId),
			blockeduser.SubjectPrimaryKey(subjectId))).
		Exec(context.TODO())
	if err != nil {
		return false, err
	}
	return true, nil
}
