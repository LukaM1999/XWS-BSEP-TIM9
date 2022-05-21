package persistence

import (
	"context"
	auth "dislinkt/common/domain"
	"dislinkt/security_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	DATABASE    = "security_service"
	COLLECTION1 = "user"
	COLLECTION2 = "rolePermission"
	COLLECTION3 = "otpSecret"
	COLLECTION4 = "userVerification"
)

type UserMongoDBStore struct {
	users             *mongo.Collection
	rolePermissions   *mongo.Collection
	otpSecrets        *mongo.Collection
	userVerifications *mongo.Collection
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION1)
	rolePermissions := client.Database(DATABASE).Collection(COLLECTION2)
	otpSecrets := client.Database(DATABASE).Collection(COLLECTION3)
	userVerifications := client.Database(DATABASE).Collection(COLLECTION4)
	return &UserMongoDBStore{
		users:             users,
		rolePermissions:   rolePermissions,
		otpSecrets:        otpSecrets,
		userVerifications: userVerifications,
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

func (store *UserMongoDBStore) Register(user *auth.User) (*auth.User, error) {
	result, err := store.users.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (store *UserMongoDBStore) Update(id primitive.ObjectID, username string) (string, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"username": username}}
	_, err := store.users.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "", err
	}
	return username, nil
}

func (store *UserMongoDBStore) Delete(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := store.users.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (store *UserMongoDBStore) DeleteAll() error {
	_, err := store.users.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	_, err = store.rolePermissions.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	_, err = store.userVerifications.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *UserMongoDBStore) CreateRolePermission(rolePermission *auth.RolePermission) (*auth.RolePermission, error) {
	_, err := store.rolePermissions.InsertOne(context.TODO(), rolePermission)
	if err != nil {
		return nil, err
	}
	return rolePermission, nil
}

func (store *UserMongoDBStore) CreateUserVerification(userVerification *domain.UserVerification) (*domain.UserVerification, error) {
	_, err := store.userVerifications.InsertOne(context.TODO(), userVerification)
	if err != nil {
		return nil, err
	}
	return userVerification, nil
}

func (store *UserMongoDBStore) SaveOTPSecret(username string, secret string) error {
	otpSecret := auth.OTPSecret{
		Username: username,
		Secret:   secret,
	}
	_, err := store.otpSecrets.InsertOne(context.TODO(), otpSecret)
	if err != nil {
		return err
	}
	return nil
}

func (store *UserMongoDBStore) GetOTPSecret(username string) (string, error) {
	filter := bson.M{"username": username}
	result := store.otpSecrets.FindOne(context.TODO(), filter)
	var otpSecret auth.OTPSecret
	err := result.Decode(&otpSecret)
	if err != nil {
		return "", err
	}
	return otpSecret.Secret, nil
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

func (store *UserMongoDBStore) getUserVerificationByToken(token string) (*domain.UserVerification, error) {
	filter := bson.M{"token": token}
	return store.filterOneVerification(filter)
}

func (store *UserMongoDBStore) filterVerification(filter interface{}) ([]*domain.UserVerification, error) {
	cursor, err := store.userVerifications.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeVerification(cursor)
}

func (store *UserMongoDBStore) filterOneVerification(filter interface{}) (UserVerification *domain.UserVerification, err error) {
	result := store.userVerifications.FindOne(context.TODO(), filter)
	err = result.Decode(&UserVerification)
	return
}

func decodeVerification(cursor *mongo.Cursor) (verifications []*domain.UserVerification, err error) {
	for cursor.Next(context.TODO()) {
		var UserVerification domain.UserVerification
		err = cursor.Decode(&UserVerification)
		if err != nil {
			return
		}
		verifications = append(verifications, &UserVerification)
	}
	err = cursor.Err()
	return
}

func (store *UserMongoDBStore) VerifyUser(token string) (string, error) {
	userVerification, err := store.filterOneVerification(bson.M{"token": token})
	if err != nil {
		return "", err
	}
	if userVerification.IsVerified {
		return "Already verified.", nil
	}
	if userVerification.TimeCreated.AddDate(0, 0, 1).Before(time.Now()) {
		return "Verification expired.", nil
	}
	_, err = store.userVerifications.UpdateOne(context.TODO(), bson.M{"token": token}, bson.M{"$set": bson.M{"isVerified": true}})
	if err != nil {
		return "", err
	}
	return "Successfully verified!", nil
}

func (store *UserMongoDBStore) UpdateUserVerification(id primitive.ObjectID, userVerification *domain.UserVerification) error {
	userVerifications, err := store.filterVerification(bson.M{"_id": id})
	if err != nil {
		return err
	}
	for _, userVerification := range userVerifications {
		_, err := store.userVerifications.UpdateOne(context.TODO(), bson.M{"_id": userVerification.Id}, bson.M{"$set": bson.M{"userVerification": userVerification}})
		if err != nil {
			return err
		}
	}
	return nil
}

func (store *UserMongoDBStore) IsVerified(username string) (bool, error) {
	userVerification, err := store.filterOneVerification(bson.M{"username": username})
	if err != nil {
		return false, err
	}
	return userVerification.IsVerified, nil
}
