package application

import (
	"bytes"
	"context"
	auth "dislinkt/common/domain"
	"dislinkt/common/tracer"
	"dislinkt/security_service/domain"
	"fmt"
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"image/png"
	"net/smtp"
	"os"
	"regexp"
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

func (service *SecurityService) Get(ctx context.Context, username string) (*auth.User, error) {
	span := tracer.StartSpanFromContext(ctx, "Get Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.Get(ctx, username)
}

func (service *SecurityService) GetAll(ctx context.Context) ([]*auth.User, error) {
	span := tracer.StartSpanFromContext(ctx, "GetAll Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.GetAll(ctx)
}

func (service *SecurityService) Register(ctx context.Context, user *auth.User, firstName string, lastName string, email string) (*auth.User, error) {
	span := tracer.StartSpanFromContext(ctx, "Register Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	registeredUser, err := service.store.Register(ctx, user)
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
		service.store.Delete(ctx, registeredUser.Id)
		return nil, err
	}
	return registeredUser, nil
}

func (service *SecurityService) CreateUserVerification(ctx context.Context, userVerification *domain.UserVerification) (*domain.UserVerification, error) {
	span := tracer.StartSpanFromContext(ctx, "CreateUserVerification Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.CreateUserVerification(ctx, userVerification)
}

func (service *SecurityService) Update(ctx context.Context, id primitive.ObjectID, username string) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "Update Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.Update(ctx, id, username)
}

func (service *SecurityService) Delete(ctx context.Context, id primitive.ObjectID) error {
	span := tracer.StartSpanFromContext(ctx, "Delete Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.Delete(ctx, id)
}

func (service *SecurityService) SetupOTP(ctx context.Context, username string) (string, []byte, error) {
	span := tracer.StartSpanFromContext(ctx, "SetupOTP Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

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

	err = service.store.SaveOTPSecret(ctx, username, key.Secret())
	if err != nil {
		return "", nil, err
	}

	return key.Secret(), buf.Bytes(), nil
}

func (service *SecurityService) GetOTPSecret(ctx context.Context, username string) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "GetOTPSecret Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.GetOTPSecret(ctx, username)
}

func (service *SecurityService) SendVerificationEmail(ctx context.Context, username string, email string, token string) error {
	span := tracer.StartSpanFromContext(ctx, "SendVerificationEmail Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

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

func (service *SecurityService) GenerateVerificationToken(ctx context.Context) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "GenerateVerificationToken Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	uuidWithHyphen := uuid.New()
	fmt.Println(uuidWithHyphen)
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	return uuid, nil
}

func (service *SecurityService) VerifyUser(ctx context.Context, token string) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "VerifyUser Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	message, err := service.store.VerifyUser(ctx, token)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return message, nil
}

func (service *SecurityService) IsVerified(ctx context.Context, username string) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "IsVerified Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	isVerified, err := service.store.IsVerified(ctx, username)
	if err != nil {
		return false, err
	}
	return isVerified, nil
}

func (service *SecurityService) SendRecoverPasswordEmail(ctx context.Context, email string, username string, token string) error {
	span := tracer.StartSpanFromContext(ctx, "SendRecoverPasswordEmail Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

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

func (service *SecurityService) UpdatePassword(ctx context.Context, token string, password string) error {
	span := tracer.StartSpanFromContext(ctx, "UpdatePassword Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	_, err := service.store.UpdatePassword(ctx, token, password)
	if err != nil {
		return err
	}
	return nil
}

func (service *SecurityService) CreatePasswordRecovery(ctx context.Context, passwordRecovery *domain.PasswordRecovery) error {
	span := tracer.StartSpanFromContext(ctx, "CreatePasswordRecovery Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.store.CreatePasswordRecovery(ctx, passwordRecovery)
	if err != nil {
		return err
	}
	return nil
}

func (service *SecurityService) GetLogs(ctx context.Context) ([]auth.Log, error) {
	span := tracer.StartSpanFromContext(ctx, "GetLogs Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	content, err := os.ReadFile(logPathPrefix + "security_service/security.log")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	logs := make([]auth.Log, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		var log auth.Log
		splitBySpace := strings.Split(line, " ")
		log.Time, err = time.Parse("2006-01-02T15:04:05.000Z", strings.Trim(strings.Split(splitBySpace[0], "=")[1], "\""))
		if err != nil {
			log.Time = time.Time{}
		}
		log.Level = strings.Split(splitBySpace[1], "=")[1]
		re := regexp.MustCompile(`msg="[/\\=!?'"\.a-zA-Z0-9_\s:-]*"`)
		msg := re.FindString(line)
		if msg != "" {
			log.Msg = strings.Trim(strings.Split(msg, "=")[1], "\"")
		}
		if msg == "" {
			re = regexp.MustCompile(`msg=[a-zA-Z]*`)
			msg = re.FindString(line)
			if msg != "" {
				log.Msg = strings.Split(msg, "=")[1]
			}
		}
		log.FullContent = line
		logs = append(logs, log)
	}
	return logs, nil
}
