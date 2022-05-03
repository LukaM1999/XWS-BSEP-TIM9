package persistence

import (
	"context"
	"dislinkt/profile_service/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "profile_service"
	COLLECTION = "profile"
)

type ProfileMongoDBStore struct {
	profiles *mongo.Collection
}

func NewProfileMongoDBStore(client *mongo.Client) domain.ProfileStore {
	profiles := client.Database(DATABASE).Collection(COLLECTION)
	return &ProfileMongoDBStore{
		profiles: profiles,
	}
}

func (store *ProfileMongoDBStore) Get(username string) (*domain.Profile, error) {
	filter := bson.M{"username": username}
	return store.filterOne(filter)
}

func (store *ProfileMongoDBStore) GetAll() ([]*domain.Profile, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *ProfileMongoDBStore) Create(profile *domain.Profile) error {
	result, err := store.profiles.InsertOne(context.TODO(), profile)
	if err != nil {
		return err
	}
	profile.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func toDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}

func (store *ProfileMongoDBStore) Update(username string, profile *domain.Profile) error {
	result, err := store.profiles.ReplaceOne(
		context.TODO(),
		bson.M{"username": username},
		profile,
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New(profile.Id.String())
	}
	return nil
}

func (store *ProfileMongoDBStore) DeleteAll() error {
	_, err := store.profiles.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *ProfileMongoDBStore) filter(filter interface{}) ([]*domain.Profile, error) {
	cursor, err := store.profiles.Find(context.TODO(), filter)
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

func (store *ProfileMongoDBStore) filterOne(filter interface{}) (profile *domain.Profile, err error) {
	result := store.profiles.FindOne(context.TODO(), filter)
	err = result.Decode(&profile)
	return
}

func decode(cursor *mongo.Cursor) (profiles []*domain.Profile, err error) {
	for cursor.Next(context.TODO()) {
		var Profile domain.Profile
		err = cursor.Decode(&Profile)
		if err != nil {
			return
		}
		profiles = append(profiles, &Profile)
	}
	err = cursor.Err()
	return
}
