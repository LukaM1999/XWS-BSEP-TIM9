package persistence

import (
	"context"
	"dislinkt/comment_service/domain"
	"dislinkt/common/tracer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "comment_service"
	COLLECTION = "comment"
)

type CommentMongoDBStore struct {
	comments *mongo.Collection
}

func NewCommentMongoDBStore(client *mongo.Client) domain.CommentStore {
	comments := client.Database(DATABASE).Collection(COLLECTION)
	return &CommentMongoDBStore{
		comments: comments,
	}
}

func (store *CommentMongoDBStore) Get(ctx context.Context, postId string) ([]*domain.Comment, error) {
	span := tracer.StartSpanFromContext(ctx, "Get Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	id, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"postId": id}
	return store.filter(filter)
}

func (store *CommentMongoDBStore) Create(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	span := tracer.StartSpanFromContext(ctx, "Create Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result, err := store.comments.InsertOne(ctx, comment)
	if err != nil {
		return nil, err
	}
	comment.Id = result.InsertedID.(primitive.ObjectID)
	return comment, nil
}

func (store *CommentMongoDBStore) DeleteAll(ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "DeleteAll Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	_, err := store.comments.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *CommentMongoDBStore) Delete(ctx context.Context, id string) error {
	span := tracer.StartSpanFromContext(ctx, "Delete Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	commentId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = store.comments.DeleteOne(ctx, bson.M{"_id": commentId})
	if err != nil {
		return err
	}
	return nil
}

func (store *CommentMongoDBStore) UpdateCommentCreator(ctx context.Context, creatorId primitive.ObjectID, creator *domain.CommentCreator) error {
	span := tracer.StartSpanFromContext(ctx, "UpdateCommentCreator Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	comments, err := store.filter(bson.M{"commentCreator._id": creatorId})
	if err != nil {
		return err
	}
	for _, comment := range comments {
		_, err := store.comments.UpdateOne(ctx, bson.M{"_id": comment.Id}, bson.M{"$set": bson.M{"commentCreator": creator}})
		if err != nil {
			return err
		}
	}
	return nil
}

func (store *CommentMongoDBStore) DeletePostComments(ctx context.Context, postId primitive.ObjectID) error {
	span := tracer.StartSpanFromContext(ctx, "DeletePostComments Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	_, err := store.comments.DeleteMany(ctx, bson.M{"postId": postId})
	if err != nil {
		return err
	}
	return nil
}

func (store *CommentMongoDBStore) filter(filter interface{}) ([]*domain.Comment, error) {
	cursor, err := store.comments.Find(context.TODO(), filter)
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

func (store *CommentMongoDBStore) filterOne(filter interface{}) (profile *domain.Comment, err error) {
	result := store.comments.FindOne(context.TODO(), filter)
	err = result.Decode(&profile)
	return
}

func decode(cursor *mongo.Cursor) (comments []*domain.Comment, err error) {
	for cursor.Next(context.TODO()) {
		var Comment domain.Comment
		err = cursor.Decode(&Comment)
		if err != nil {
			return
		}
		comments = append(comments, &Comment)
	}
	err = cursor.Err()
	return
}
