package domain

import (
	"context"
	auth "dislinkt/common/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStore interface {
	Get(ctx context.Context, username string) (*auth.User, error)
	GetAll(ctx context.Context) ([]*auth.User, error)
	Register(ctx context.Context, user *auth.User) (*auth.User, error)
	Update(ctx context.Context, id primitive.ObjectID, username string) (string, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
	DeleteAll(ctx context.Context) error
	CreateRolePermission(ctx context.Context, rolePermission *auth.RolePermission) (*auth.RolePermission, error)
	SaveOTPSecret(ctx context.Context, username string, secret string) error
	GetOTPSecret(ctx context.Context, username string) (string, error)
	CreateUserVerification(ctx context.Context, userVerification *UserVerification) (*UserVerification, error)
	VerifyUser(ctx context.Context, token string) (string, error)
	IsVerified(ctx context.Context, username string) (bool, error)
	UpdatePassword(ctx context.Context, token string, password string) (string, error)
	CreatePasswordRecovery(ctx context.Context, passwordRecovery *PasswordRecovery) error
}
