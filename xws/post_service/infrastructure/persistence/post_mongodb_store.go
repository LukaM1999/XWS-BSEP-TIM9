package persistence

import (
	"context"
	"dislinkt/post_service/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	DATABASE    = "post_service"
	COLLECTION1 = "connection"
	COLLECTION2 = "post"
)

type PostMongoDBStore struct {
	posts       *mongo.Collection
	connections *mongo.Collection
}

func NewPostMongoDBStore(client *mongo.Client) domain.PostStore {
	connections := client.Database(DATABASE).Collection(COLLECTION1)
	posts := client.Database(DATABASE).Collection(COLLECTION2)
	index := mongo.IndexModel{
		Keys:    bson.D{{"fullName", "text"}},
		Options: options.Index().SetUnique(true),
	}
	opts := options.CreateIndexes().SetMaxTime(20 * time.Second)
	_, err := posts.Indexes().CreateOne(context.TODO(), index, opts)

	if err != nil {
		panic(err)
	}

	return &PostMongoDBStore{
		connections: connections,
		posts:       posts,
	}
}

func (store *PostMongoDBStore) Get(id string) (*domain.Post, error) {
	postId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": postId}
	return store.filterOne(filter)
}

func (store *PostMongoDBStore) GetProfilePosts(profileId string) ([]*domain.Post, error) {
	id, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"profile._id": id}
	return store.filter(filter)
}

func (store *PostMongoDBStore) GetConnectionPosts(profileId string) ([]*domain.Post, error) {
	id, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"$or", bson.A{bson.M{"_issuerId": id}, bson.M{"_subjectId": id}}}, {"isApproved", true}}
	connections, err := store.filterConnections(filter)
	if err != nil {
		return nil, err
	}
	posts := make([]*domain.Post, 0)
	for _, connection := range connections {
		if connection.IssuerId == id {
			connectionPosts, err := store.filter(bson.M{"profile._id": connection.SubjectId})
			if err != nil {
				return nil, err
			}
			posts = append(posts, connectionPosts...)
		} else if connection.SubjectId == id {
			connectionPosts, err := store.filter(bson.M{"profile._id": connection.IssuerId})
			if err != nil {
				return nil, err
			}
			posts = append(posts, connectionPosts...)
		}
	}
	return posts, nil
}

func (store *PostMongoDBStore) Create(post *domain.Post) error {
	result, err := store.posts.InsertOne(context.TODO(), post)
	if err != nil {
		return err
	}
	post.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *PostMongoDBStore) CreateConnection(connection *domain.Connection) error {
	result, err := store.connections.InsertOne(context.TODO(), connection)
	if err != nil {
		return err
	}
	connection.Id = result.InsertedID.(primitive.ObjectID)
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

func (store *PostMongoDBStore) Update(id string, post *domain.Post) error {
	postId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := store.posts.ReplaceOne(
		context.TODO(),
		bson.M{"_id": postId},
		post,
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New(post.Id.String())
	}
	return nil
}

func (store *PostMongoDBStore) UpdateProfile(id primitive.ObjectID, profile *domain.Profile) error {
	posts, err := store.filter(bson.M{"profile._id": id})
	if err != nil {
		return err
	}
	for _, post := range posts {
		_, err := store.posts.UpdateOne(context.TODO(), bson.M{"_id": post.Id}, bson.M{"$set": bson.M{"profile": profile}})
		if err != nil {
			return err
		}
	}
	return nil
}

func (store *PostMongoDBStore) Delete(id string) error {
	postId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = store.posts.DeleteOne(context.TODO(), bson.M{"_id": postId})
	if err != nil {
		return err
	}
	return nil
}

func (store *PostMongoDBStore) DeleteAll() error {
	_, err := store.posts.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	_, err = store.connections.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *PostMongoDBStore) filter(filter interface{}) ([]*domain.Post, error) {
	cursor, err := store.posts.Find(context.TODO(), filter)
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

func (store *PostMongoDBStore) filterConnections(filter interface{}) ([]*domain.Connection, error) {
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
	return decodeConnections(cursor)
}

func (store *PostMongoDBStore) filterOne(filter interface{}) (post *domain.Post, err error) {
	result := store.posts.FindOne(context.TODO(), filter)
	err = result.Decode(&post)
	return
}

func decode(cursor *mongo.Cursor) (posts []*domain.Post, err error) {
	for cursor.Next(context.TODO()) {
		var Post domain.Post
		err = cursor.Decode(&Post)
		if err != nil {
			return
		}
		posts = append(posts, &Post)
	}
	err = cursor.Err()
	return
}

func decodeConnections(cursor *mongo.Cursor) (connections []*domain.Connection, err error) {
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
