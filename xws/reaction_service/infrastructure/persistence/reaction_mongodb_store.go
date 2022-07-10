package persistence

import (
	"context"
	"dislinkt/common/tracer"
	"dislinkt/reaction_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE   = "reaction_service"
	COLLECTION = "reaction"
)

type ReactionMongoDBStore struct {
	reactions *mongo.Collection
}

func NewReactionMongoDBStore(client *mongo.Client) domain.ReactionStore {
	reactions := client.Database(DATABASE).Collection(COLLECTION)
	return &ReactionMongoDBStore{
		reactions: reactions,
	}
}

func (store *ReactionMongoDBStore) Get(ctx context.Context, postId string) ([]*domain.Reaction, error) {
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

func (store *ReactionMongoDBStore) Reaction(ctx context.Context, reaction *domain.Reaction) (*domain.Reaction, error) {
	span := tracer.StartSpanFromContext(ctx, "Reaction Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"_id": reaction.Id}
	existingReaction, err := store.filter(filter)
	if err != nil {
		return nil, err
	}
	if len(existingReaction) > 0 {
		result, err2 := store.UpdateReaction(ctx, reaction, filter)
		if err2 != nil {
			return nil, err2
		}
		return result, nil
	} else {
		result, err := store.reactions.InsertOne(ctx, reaction)
		if err != nil {
			return nil, err
		}
		reaction.Id = result.InsertedID.(primitive.ObjectID)
	}
	return reaction, nil
}

func (store *ReactionMongoDBStore) UpdateReaction(ctx context.Context, reaction *domain.Reaction, filter bson.M) (updatedReaction *domain.Reaction, err error) {
	span := tracer.StartSpanFromContext(ctx, "UpdateReaction Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	update := bson.M{"$set": bson.M{"createdAt": reaction.CreatedAt, "type": reaction.Type}}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	result := store.reactions.FindOneAndUpdate(ctx, filter, update, &opt)
	if result.Err() != nil {
		return nil, result.Err()
	}
	err = result.Decode(&updatedReaction)
	return
}

func (store *ReactionMongoDBStore) DeleteAll(ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "DeleteAll Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	_, err := store.reactions.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *ReactionMongoDBStore) Delete(ctx context.Context, id string) error {
	span := tracer.StartSpanFromContext(ctx, "Delete Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	reactionId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = store.reactions.DeleteOne(ctx, bson.M{"_id": reactionId})
	if err != nil {
		return err
	}
	return nil
}

func (store *ReactionMongoDBStore) DeletePostReactions(ctx context.Context, postId primitive.ObjectID) error {
	span := tracer.StartSpanFromContext(ctx, "DeletePostReactions Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	_, err := store.reactions.DeleteMany(ctx, bson.M{"postId": postId})
	if err != nil {
		return err
	}
	return nil
}

func (store *ReactionMongoDBStore) filter(filter interface{}) ([]*domain.Reaction, error) {
	cursor, err := store.reactions.Find(context.TODO(), filter)
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

func (store *ReactionMongoDBStore) filterOne(filter interface{}) (reaction *domain.Reaction, err error) {
	result := store.reactions.FindOne(context.TODO(), filter)
	err = result.Decode(&reaction)
	return
}

func decode(cursor *mongo.Cursor) (reactions []*domain.Reaction, err error) {
	for cursor.Next(context.TODO()) {
		var Reaction domain.Reaction
		err = cursor.Decode(&Reaction)
		if err != nil {
			return
		}
		reactions = append(reactions, &Reaction)
	}
	err = cursor.Err()
	return
}
