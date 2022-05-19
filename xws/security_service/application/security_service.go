package application

import (
	"bytes"
	auth "dislinkt/common/domain"
	"dislinkt/security_service/domain"
	"github.com/pquerna/otp/totp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"image/png"
)

type SecurityService struct {
	store domain.UserStore
}

func NewSecurityService(store domain.UserStore) *SecurityService {
	return &SecurityService{
		store: store,
	}
}

func (service *SecurityService) Get(username string) (*auth.User, error) {
	return service.store.Get(username)
}

func (service *SecurityService) GetAll() ([]*auth.User, error) {
	return service.store.GetAll()
}

func (service *SecurityService) Register(user *auth.User) (*auth.User, error) {
	return service.store.Register(user)
}

func (service *SecurityService) Update(id primitive.ObjectID, username string) (string, error) {
	return service.store.Update(id, username)
}

func (service *SecurityService) Delete(id primitive.ObjectID) error {
	return service.store.Delete(id)
}

func (service *SecurityService) SetupOTP(username string) (string, []byte, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "dislinkt.com",
		AccountName: username,
	})
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		panic(err)
	}
	png.Encode(&buf, img)

	err = service.store.SaveOTPSecret(username, key.Secret())
	if err != nil {
		return "", nil, err
	}

	return key.Secret(), buf.Bytes(), nil
}

func (service *SecurityService) GetOTPSecret(username string) (string, error) {
	return service.store.GetOTPSecret(username)
}
