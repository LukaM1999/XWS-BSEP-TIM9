package persistence

import (
	"context"
	"dislinkt/common/tracer"
	"dislinkt/post_service/domain"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE    = "post_service"
	COLLECTION1 = "connection"
	COLLECTION2 = "post"
	COLLECTION3 = "job"
)

type PostMongoDBStore struct {
	posts       *mongo.Collection
	connections *mongo.Collection
	jobs        *mongo.Collection
}

func NewPostMongoDBStore(client *mongo.Client) domain.PostStore {
	connections := client.Database(DATABASE).Collection(COLLECTION1)
	posts := client.Database(DATABASE).Collection(COLLECTION2)
	jobs := client.Database(DATABASE).Collection(COLLECTION3)
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
		jobs:        jobs,
	}
}

func (store *PostMongoDBStore) UpdatePostImage(ctx context.Context, id primitive.ObjectID, url string) (*domain.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "UpdatePostImage Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	_, err := store.posts.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"content.image": url}})
	if err != nil {
		return nil, err
	}
	post, err := store.filterOne(bson.M{"_id": id})
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (store *PostMongoDBStore) Get(ctx context.Context, id string) (*domain.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "Get Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	postId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": postId}
	return store.filterOne(filter)
}

func (store *PostMongoDBStore) GetProfilePosts(ctx context.Context, profileId string) ([]*domain.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "GetProfilePosts Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	id, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"profile._id": id}
	return store.filter(filter)
}

func (store *PostMongoDBStore) GetConnectionPosts(ctx context.Context, profileId string) ([]*domain.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "GetConnectionPosts Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	id, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"$or", bson.A{bson.M{"_issuerId": id}, bson.M{"_subjectId": id}}}}
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

func (store *PostMongoDBStore) Create(ctx context.Context, post *domain.Post) error {
	span := tracer.StartSpanFromContext(ctx, "Create Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result, err := store.posts.InsertOne(ctx, post)
	if err != nil {
		return err
	}
	post.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *PostMongoDBStore) CreateConnection(ctx context.Context, connection *domain.Connection) error {
	span := tracer.StartSpanFromContext(ctx, "CreateConnection Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result, err := store.connections.InsertOne(ctx, connection)
	if err != nil {
		return err
	}
	connection.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *PostMongoDBStore) DeleteConnection(ctx context.Context, id primitive.ObjectID) error {
	span := tracer.StartSpanFromContext(ctx, "DeleteConnection Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	_, err := store.connections.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (store *PostMongoDBStore) Update(ctx context.Context, id string, post *domain.Post) error {
	span := tracer.StartSpanFromContext(ctx, "Update Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	postId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := store.posts.ReplaceOne(
		ctx,
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

func (store *PostMongoDBStore) UpdateProfile(ctx context.Context, id primitive.ObjectID, profile *domain.Profile) error {
	span := tracer.StartSpanFromContext(ctx, "UpdateProfile Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts, err := store.filter(bson.M{"profile._id": id})
	if err != nil {
		return err
	}
	for _, post := range posts {
		_, err := store.posts.UpdateOne(ctx, bson.M{"_id": post.Id}, bson.M{"$set": bson.M{"profile": profile}})
		if err != nil {
			return err
		}
	}
	return nil
}

func (store *PostMongoDBStore) Delete(ctx context.Context, id string) error {
	span := tracer.StartSpanFromContext(ctx, "Delete Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	postId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = store.posts.DeleteOne(ctx, bson.M{"_id": postId})
	if err != nil {
		return err
	}
	return nil
}

func (store *PostMongoDBStore) DeleteAll(ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "DeleteAll Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	_, err := store.posts.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	_, err = store.connections.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	_, err = store.jobs.DeleteMany(ctx, bson.D{{}})
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
