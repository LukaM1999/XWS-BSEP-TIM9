package application

import (
	"bytes"
	auth "dislinkt/common/domain"
	"dislinkt/security_service/domain"
	"fmt"
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"image/png"
	"net/smtp"
	"os"
	"strings"
	"text/template"
	"time"
)

type SecurityService struct {
	store        domain.UserStore
	orchestrator *CreateProfileOrchestrator
}

func NewSecurityService(store domain.UserStore, orchestrator *CreateProfileOrchestrator) *SecurityService {
	return &SecurityService{
		store:        store,
		orchestrator: orchestrator,
	}
}

func (service *SecurityService) Get(username string) (*auth.User, error) {
	return service.store.Get(username)
}

func (service *SecurityService) GetAll() ([]*auth.User, error) {
	return service.store.GetAll()
}

func (service *SecurityService) Register(user *auth.User, firstName string, lastName string, email string) (*auth.User, error) {
	registeredUser, err := service.store.Register(user)
	if err != nil {
		return nil, err
	}
	profile := &auth.Profile{
		Id:             registeredUser.Id,
		Username:       registeredUser.Username,
		FirstName:      firstName,
		LastName:       lastName,
		FullName:       firstName + " " + lastName,
		DateOfBirth:    time.Time{},
		PhoneNumber:    "",
		Email:          email,
		Gender:         "",
		IsPrivate:      false,
		Biography:      "",
		Education:      nil,
		WorkExperience: nil,
		Skills:         nil,
		Interests:      nil,
		AgentToken:     "",
	}
	err = service.orchestrator.Start(profile)
	if err != nil {
		service.store.Delete(registeredUser.Id)
		return nil, err
	}
	return registeredUser, nil
}

func (service *SecurityService) CreateUserVerification(userVerification *domain.UserVerification) (*domain.UserVerification, error) {
	return service.store.CreateUserVerification(userVerification)
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

func (service *SecurityService) SendVerificationEmail(username string, email string, token string) error {
	// Sender data.
	from := "isatestmail2021@gmail.com"
	password := "yciuowcxhvykcots"

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, err := template.ParseFiles("./templates/template.html")
	if err != nil {
		fmt.Println(err)
		return err
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Name    string
		Message string
		Token   string
	}{
		Name:    username,
		Message: "Click link below to verify your account",
		Token:   token,
	})

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent!")
	return nil
}

func (service *SecurityService) GenerateVerificationToken() (string, error) {
	uuidWithHyphen := uuid.New()
	fmt.Println(uuidWithHyphen)
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	return uuid, nil
}

func (service *SecurityService) VerifyUser(token string) (string, error) {
	message, err := service.store.VerifyUser(token)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return message, nil
}

func (service *SecurityService) IsVerified(username string) (bool, error) {
	isVerified, err := service.store.IsVerified(username)
	if err != nil {
		return false, err
	}
	return isVerified, nil
}

func (service *SecurityService) SendRecoverPasswordEmail(email string, username string, token string) error {
	// Sender data.
	from := "isatestmail2021@gmail.com"
	password := "yciuowcxhvykcots"

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, err := template.ParseFiles("./templates/recoverPassword.html")
	if err != nil {
		fmt.Println(err)
		return err
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Name    string
		Message string
		Token   string
	}{
		Name:    username,
		Message: "Click link below to recover your password",
		Token:   token,
	})

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent!")

	return nil
}

func (service *SecurityService) UpdatePassword(token string, password string) error {
	_, err := service.store.UpdatePassword(token, password)
	if err != nil {
		return err
	}
	return nil
}

func (service *SecurityService) CreatePasswordRecovery(passwordRecovery *domain.PasswordRecovery) error {
	err := service.store.CreatePasswordRecovery(passwordRecovery)
	if err != nil {
		return err
	}
	return nil
}

func (service *SecurityService) GetLogs() ([]string, error) {
	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	content, err := os.ReadFile(logPathPrefix + "security_service/security.log")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}
