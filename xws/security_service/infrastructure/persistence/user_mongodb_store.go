package persistence

import (
	"context"
	auth "dislinkt/common/domain"
	"dislinkt/security_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "security_service"
	COLLECTION = "user"
)

type UserMongoDBStore struct {
	users *mongo.Collection
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &UserMongoDBStore{
		users: users,
	}
}

func (store *UserMongoDBStore) Get(username string) (*auth.User, error) {
	filter := bson.M{"username": username}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) GetAll() ([]*auth.User, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *UserMongoDBStore) Register(User *auth.User) error {
	result, err := store.users.InsertOne(context.TODO(), User)
	if err != nil {
		return err
	}
	User.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *UserMongoDBStore) DeleteAll() error {
	_, err := store.users.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *UserMongoDBStore) filter(filter interface{}) ([]*auth.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
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

func (store *UserMongoDBStore) filterOne(filter interface{}) (User *auth.User, err error) {
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&User)
	return
}

func decode(cursor *mongo.Cursor) (users []*auth.User, err error) {
	for cursor.Next(context.TODO()) {
		var User auth.User
		err = cursor.Decode(&User)
		if err != nil {
			return
		}
		users = append(users, &User)
	}
	err = cursor.Err()
	return
}
