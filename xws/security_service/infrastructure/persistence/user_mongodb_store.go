package persistence

import (
	"context"
	auth "dislinkt/common/domain"
	"dislinkt/common/tracer"
	"dislinkt/security_service/domain"
	"dislinkt/security_service/infrastructure/api"
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
	COLLECTION5 = "passwordRecovery"
)

type UserMongoDBStore struct {
	users              *mongo.Collection
	rolePermissions    *mongo.Collection
	otpSecrets         *mongo.Collection
	userVerifications  *mongo.Collection
	passwordRecoveries *mongo.Collection
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION1)
	rolePermissions := client.Database(DATABASE).Collection(COLLECTION2)
	otpSecrets := client.Database(DATABASE).Collection(COLLECTION3)
	userVerifications := client.Database(DATABASE).Collection(COLLECTION4)
	passwordRecoveries := client.Database(DATABASE).Collection(COLLECTION5)
	return &UserMongoDBStore{
		users:              users,
		rolePermissions:    rolePermissions,
		otpSecrets:         otpSecrets,
		userVerifications:  userVerifications,
		passwordRecoveries: passwordRecoveries,
	}
}

func (store *UserMongoDBStore) Get(ctx context.Context, username string) (*auth.User, error) {
	span := tracer.StartSpanFromContext(ctx, "Get Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"username": username}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) GetAll(ctx context.Context) ([]*auth.User, error) {
	span := tracer.StartSpanFromContext(ctx, "GetAll Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *UserMongoDBStore) Register(ctx context.Context, user *auth.User) (*auth.User, error) {
	span := tracer.StartSpanFromContext(ctx, "Register Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result, err := store.users.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (store *UserMongoDBStore) Update(ctx context.Context, id primitive.ObjectID, username string) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "Update Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"username": username}}
	_, err := store.users.UpdateOne(ctx, filter, update)
	if err != nil {
		return "", err
	}
	return username, nil
}

func (store *UserMongoDBStore) Delete(ctx context.Context, id primitive.ObjectID) error {
	span := tracer.StartSpanFromContext(ctx, "Delete Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"_id": id}
	_, err := store.users.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (store *UserMongoDBStore) DeleteAll(ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "DeleteAll Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	_, err := store.users.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	_, err = store.rolePermissions.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	_, err = store.userVerifications.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	_, err = store.passwordRecoveries.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *UserMongoDBStore) CreateRolePermission(ctx context.Context, rolePermission *auth.RolePermission) (*auth.RolePermission, error) {
	span := tracer.StartSpanFromContext(ctx, "CreateRolePermission Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	_, err := store.rolePermissions.InsertOne(ctx, rolePermission)
	if err != nil {
		return nil, err
	}
	return rolePermission, nil
}

func (store *UserMongoDBStore) CreatePasswordRecovery(ctx context.Context, passwordRecovery *domain.PasswordRecovery) error {
	span := tracer.StartSpanFromContext(ctx, "CreatePasswordRecovery Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	_, err := store.passwordRecoveries.InsertOne(ctx, passwordRecovery)
	if err != nil {
		return err
	}
	return nil
}

func (store *UserMongoDBStore) CreateUserVerification(ctx context.Context, userVerification *domain.UserVerification) (*domain.UserVerification, error) {
	span := tracer.StartSpanFromContext(ctx, "CreateUserVerification Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	_, err := store.userVerifications.InsertOne(ctx, userVerification)
	if err != nil {
		return nil, err
	}
	return userVerification, nil
}

func (store *UserMongoDBStore) SaveOTPSecret(ctx context.Context, username string, secret string) error {
	span := tracer.StartSpanFromContext(ctx, "SaveOTPSecret Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	otpSecret := auth.OTPSecret{
		Username: username,
		Secret:   secret,
	}
	_, err := store.otpSecrets.InsertOne(ctx, otpSecret)
	if err != nil {
		return err
	}
	return nil
}

func (store *UserMongoDBStore) GetOTPSecret(ctx context.Context, username string) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "GetOTPSecret Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"username": username}
	result := store.otpSecrets.FindOne(ctx, filter)
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

func (store *UserMongoDBStore) VerifyUser(ctx context.Context, token string) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "VerifyUser Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

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
	_, err = store.userVerifications.UpdateOne(ctx, bson.M{"token": token}, bson.M{"$set": bson.M{"isVerified": true}})
	if err != nil {
		return "", err
	}
	return "Successfully verified!", nil
}

func (store *UserMongoDBStore) UpdateUserVerification(ctx context.Context, id primitive.ObjectID) error {
	span := tracer.StartSpanFromContext(ctx, "UpdateUserVerification Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	userVerifications, err := store.filterVerification(bson.M{"_id": id})
	if err != nil {
		return err
	}
	for _, userVerification := range userVerifications {
		_, err := store.userVerifications.UpdateOne(ctx, bson.M{"_id": userVerification.Id}, bson.M{"$set": bson.M{"userVerification": userVerification}})
		if err != nil {
			return err
		}
	}
	return nil
}

func (store *UserMongoDBStore) IsVerified(ctx context.Context, username string) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "IsVerified Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	userVerification, err := store.filterOneVerification(bson.M{"username": username})
	if err != nil {
		return false, err
	}
	return userVerification.IsVerified, nil
}

func (store *UserMongoDBStore) UpdatePassword(ctx context.Context, token string, password string) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "UpdatePassword Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	passwordRecovery, err := store.filterOnePasswordRecovery(bson.M{"token": token})
	if err != nil {
		return "", err
	}
	if passwordRecovery.IsRecovered {
		return "Already recovered", nil
	}
	if passwordRecovery.TimeCreated.AddDate(0, 0, 1).Before(time.Now()) {
		return "Recovery expired.", nil
	}
	_, err = store.passwordRecoveries.UpdateOne(ctx, bson.M{"token": token}, bson.M{"$set": bson.M{"isRecovered": true}})
	if err != nil {
		return "", err
	}
	_, err = store.users.UpdateOne(ctx, bson.M{"username": passwordRecovery.Username},
		bson.M{"$set": bson.M{"password": api.HashPassword(password)}})
	if err != nil {
		return "", err
	}
	_, err = store.userVerifications.UpdateOne(ctx, bson.M{"username": passwordRecovery.Username},
		bson.M{"$set": bson.M{"isVerified": true}})
	if err != nil {
		return "", err
	}
	return "Password is recovered", nil
}

func (store *UserMongoDBStore) filterOnePasswordRecovery(filter interface{}) (PasswordRecovery *domain.PasswordRecovery, err error) {
	result := store.passwordRecoveries.FindOne(context.TODO(), filter)
	err = result.Decode(&PasswordRecovery)
	return
}
