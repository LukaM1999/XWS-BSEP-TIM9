package persistence

import (
	"context"
	"dislinkt/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE    = "connection_service"
	COLLECTION1 = "connection"
	COLLECTION2 = "profilePrivacy"
)

type ConnectionMongoDBStore struct {
	connections     *mongo.Collection
	profilesPrivacy *mongo.Collection
}

func NewConnectionMongoDBStore(client *mongo.Client) domain.ConnectionStore {
	connections := client.Database(DATABASE).Collection(COLLECTION1)
	profilesPrivacy := client.Database(DATABASE).Collection(COLLECTION2)
	return &ConnectionMongoDBStore{
		connections:     connections,
		profilesPrivacy: profilesPrivacy,
	}
}

func (store *ConnectionMongoDBStore) Get(userId string) ([]*domain.Connection, error) {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"$or": []bson.M{{"subjectId": id},
		{"issuerId": id}}}
	return store.filter(filter)
}

func (store *ConnectionMongoDBStore) Create(connection *domain.Connection) (*domain.Connection, error) {
	//filter := bson.M{"userId": connection.SubjectId}
	//privacy, err := store.filterOnePrivacy(filter)
	//if err != nil {
	//	return nil, err
	//}
	connection.IsApproved = true
	result, err := store.connections.InsertOne(context.TODO(), connection)
	if err != nil {
		return nil, err
	}
	connection.Id = result.InsertedID.(primitive.ObjectID)
	return connection, nil
}

func (store *ConnectionMongoDBStore) CreatePrivacy(privacy *domain.ProfilePrivacy) (*domain.ProfilePrivacy, error) {
	result, err := store.profilesPrivacy.InsertOne(context.TODO(), privacy)
	if err != nil {
		return nil, err
	}
	privacy.Id = result.InsertedID.(primitive.ObjectID)
	return privacy, nil
}

func (store *ConnectionMongoDBStore) DeleteAll() error {
	_, err := store.connections.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	_, err = store.profilesPrivacy.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectionMongoDBStore) Delete(id string) error {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": Id}
	_, err = store.connections.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectionMongoDBStore) Update(id string) (*domain.Connection, error) {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": Id}
	connection, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}
	connection.IsApproved = !connection.IsApproved
	_, err = store.connections.UpdateOne(context.TODO(), filter, bson.D{{"$set", bson.M{"isApproved": connection.IsApproved}}})
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func (store *ConnectionMongoDBStore) UpdatePrivacy(id primitive.ObjectID) error {
	filter := bson.M{"userId": id}
	privacy, err := store.filterOnePrivacy(filter)
	if err != nil {
		return err
	}
	privacy.IsPrivate = !privacy.IsPrivate
	_, err = store.profilesPrivacy.UpdateOne(context.TODO(), filter, bson.D{{"$set", bson.M{"isPrivate": privacy.IsPrivate}}})
	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectionMongoDBStore) CreateProfilePrivacy(privacy *domain.ProfilePrivacy) (*domain.ProfilePrivacy, error) {
	result, err := store.profilesPrivacy.InsertOne(context.TODO(), privacy)
	if err != nil {
		return nil, err
	}
	privacy.Id = result.InsertedID.(primitive.ObjectID)
	return privacy, nil
}

func (store *ConnectionMongoDBStore) DeleteProfilePrivacy(id primitive.ObjectID) error {
	filter := bson.M{"userId": id}
	_, err := store.profilesPrivacy.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectionMongoDBStore) filter(filter interface{}) ([]*domain.Connection, error) {
	cursor, err := store.connections.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *ConnectionMongoDBStore) filterOne(filter interface{}) (connection *domain.Connection, err error) {
	result := store.connections.FindOne(context.TODO(), filter)
	err = result.Decode(&connection)
	return
}

func (store *ConnectionMongoDBStore) filterOnePrivacy(filter interface{}) (privacy *domain.ProfilePrivacy, err error) {
	result := store.profilesPrivacy.FindOne(context.TODO(), filter)
	err = result.Decode(&privacy)
	return
}

func decode(cursor *mongo.Cursor) (connections []*domain.Connection, err error) {
	for cursor.Next(context.TODO()) {
		var Connection domain.Connection
		err = cursor.Decode(&Connection)
		if err != nil {
			return
		}
		connections = append(connections, &Connection)
	}
	err = cursor.Err()
	return
}
