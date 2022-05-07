package persistence

import (
	"context"
	"dislinkt/comment_service/domain"
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

func (store *CommentMongoDBStore) Get(postId string) ([]*domain.Comment, error) {
	id, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"postId": id}
	return store.filter(filter)
}

func (store *CommentMongoDBStore) Create(comment *domain.Comment) (*domain.Comment, error) {
	result, err := store.comments.InsertOne(context.TODO(), comment)
	if err != nil {
		return nil, err
	}
	comment.Id = result.InsertedID.(primitive.ObjectID)
	return comment, nil
}

func (store *CommentMongoDBStore) DeleteAll() error {
	_, err := store.comments.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *CommentMongoDBStore) Delete(id string) error {
	commentId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = store.comments.DeleteOne(context.TODO(), bson.M{"_id": commentId})
	if err != nil {
		return err
	}
	return nil
}

func (store *CommentMongoDBStore) UpdateCommentCreator(creatorId primitive.ObjectID, creator *domain.CommentCreator) error {
	comments, err := store.filter(bson.M{"commentCreator._id": creatorId})
	if err != nil {
		return err
	}
	for _, comment := range comments {
		_, err := store.comments.UpdateOne(context.TODO(), bson.M{"_id": comment.Id}, bson.M{"$set": bson.M{"commentCreator": creator}})
		if err != nil {
			return err
		}
	}
	return nil
}

func (store *CommentMongoDBStore) DeletePostComments(postId primitive.ObjectID) error {
	_, err := store.comments.DeleteMany(context.TODO(), bson.M{"postId": postId})
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
