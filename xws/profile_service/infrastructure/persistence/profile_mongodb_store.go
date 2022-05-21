package persistence

import (
	"context"
	"dislinkt/profile_service/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
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
	_, err := profiles.Indexes().DropOne(context.TODO(), "fullName_text")
	index := mongo.IndexModel{
		Keys: bson.D{{"fullName", "text"}},
		//Options: options.Index().SetUnique(true),
	}
	opts := options.CreateIndexes().SetMaxTime(20 * time.Second)
	_, err = profiles.Indexes().CreateOne(context.TODO(), index, opts)

	if err != nil {
		panic(err)
	}

	return &ProfileMongoDBStore{
		profiles: profiles,
	}
}

func (store *ProfileMongoDBStore) Get(profileId string) (*domain.Profile, error) {
	id, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *ProfileMongoDBStore) GetAll(search string) ([]*domain.Profile, error) {
	filter := bson.D{{"fullName", bson.M{"$regex": "^.*" + search + ".*$"}}, {"isPrivate", false}}
	return store.filter(filter, search)
}

func (store *ProfileMongoDBStore) Create(profile *domain.Profile) error {
	result, err := store.profiles.InsertOne(context.TODO(), profile)
	if err != nil {
		return err
	}
	profile.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *ProfileMongoDBStore) Update(profileId string, profile *domain.Profile) error {
	id, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return err
	}
	result, err := store.profiles.ReplaceOne(
		context.TODO(),
		bson.M{"_id": id},
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

func (store *ProfileMongoDBStore) filter(filter interface{}, search string) ([]*domain.Profile, error) {
	cursor, err := store.profiles.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, errors.New(search)
	}
	return decode(cursor)
}

func (store *ProfileMongoDBStore) filterOne(filter interface{}) (profile *domain.Profile, err error) {
	result := store.profiles.FindOne(context.TODO(), filter)
	err = result.Decode(&profile)
	return
}

func (store *ProfileMongoDBStore) Delete(id string) error {
	profileId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = store.profiles.DeleteOne(context.TODO(), bson.M{"_id": profileId})
	if err != nil {
		return err
	}
	return nil
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
