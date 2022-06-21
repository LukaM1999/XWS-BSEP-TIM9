package config

import (
	"dislinkt/common/loggers"
	"github.com/joho/godotenv"
	"os"
)

var log = loggers.NewSecurityLogger()

type Config struct {
	Port                        string
	SecurityDBHost              string
	SecurityDBPort              string
	ProfileHost                 string
	ProfilePort                 string
	NatsHost                    string
	NatsPort                    string
	NatsUser                    string
	NatsPass                    string
	CreateProfileCommandSubject string
	CreateProfileReplySubject   string
	UpdateProfileCommandSubject string
	UpdateProfileReplySubject   string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:                        os.Getenv("SECURITY_SERVICE_PORT"),
		SecurityDBHost:              os.Getenv("SECURITY_DB_HOST"),
		SecurityDBPort:              os.Getenv("SECURITY_DB_PORT"),
		ProfileHost:                 os.Getenv("PROFILE_SERVICE_HOST"),
		ProfilePort:                 os.Getenv("PROFILE_SERVICE_PORT"),
		NatsHost:                    os.Getenv("NATS_HOST"),
		NatsPort:                    os.Getenv("NATS_PORT"),
		NatsUser:                    os.Getenv("NATS_USER"),
		NatsPass:                    os.Getenv("NATS_PASS"),
		CreateProfileCommandSubject: os.Getenv("CREATE_PROFILE_COMMAND_SUBJECT"),
		CreateProfileReplySubject:   os.Getenv("CREATE_PROFILE_REPLY_SUBJECT"),
		UpdateProfileCommandSubject: os.Getenv("UPDATE_PROFILE_COMMAND_SUBJECT"),
		UpdateProfileReplySubject:   os.Getenv("UPDATE_PROFILE_REPLY_SUBJECT"),
	}
}

func SetEnvironment() error {
	if os.Getenv("OS_ENV") != "docker" {
		if err := godotenv.Load("../.env.dev"); err != nil {
			log.Fatal("ENVF")
		}
	}
	return nil
}
