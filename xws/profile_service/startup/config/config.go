package config

import (
	"dislinkt/common/loggers"
	"github.com/joho/godotenv"
	"os"
)

var log = loggers.NewProfileLogger()

type Config struct {
	Port                        string
	ProfileDBHost               string
	ProfileDBPort               string
	PostHost                    string
	PostPort                    string
	CommentHost                 string
	CommentPort                 string
	SecurityHost                string
	SecurityPort                string
	NatsHost                    string
	NatsPort                    string
	NatsUser                    string
	NatsPass                    string
	CreateProfileCommandSubject string
	CreateProfileReplySubject   string
	UpdateProfileCommandSubject string
	UpdateProfileReplySubject   string
	PromoteJobCommandSubject    string
	PromoteJobReplySubject      string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:                        os.Getenv("PROFILE_SERVICE_PORT"),
		ProfileDBHost:               os.Getenv("PROFILE_DB_HOST"),
		ProfileDBPort:               os.Getenv("PROFILE_DB_PORT"),
		PostHost:                    os.Getenv("POST_SERVICE_HOST"),
		PostPort:                    os.Getenv("POST_SERVICE_PORT"),
		CommentHost:                 os.Getenv("COMMENT_SERVICE_HOST"),
		CommentPort:                 os.Getenv("COMMENT_SERVICE_PORT"),
		SecurityHost:                os.Getenv("SECURITY_SERVICE_HOST"),
		SecurityPort:                os.Getenv("SECURITY_SERVICE_PORT"),
		NatsHost:                    os.Getenv("NATS_HOST"),
		NatsPort:                    os.Getenv("NATS_PORT"),
		NatsUser:                    os.Getenv("NATS_USER"),
		NatsPass:                    os.Getenv("NATS_PASS"),
		CreateProfileCommandSubject: os.Getenv("CREATE_PROFILE_COMMAND_SUBJECT"),
		CreateProfileReplySubject:   os.Getenv("CREATE_PROFILE_REPLY_SUBJECT"),
		UpdateProfileCommandSubject: os.Getenv("UPDATE_PROFILE_COMMAND_SUBJECT"),
		UpdateProfileReplySubject:   os.Getenv("UPDATE_PROFILE_REPLY_SUBJECT"),
		PromoteJobCommandSubject:    os.Getenv("PROMOTE_JOB_COMMAND_SUBJECT"),
		PromoteJobReplySubject:      os.Getenv("PROMOTE_JOB_REPLY_SUBJECT"),
	}
}

func SetEnvironment() error {
	if os.Getenv("OS_ENV") != "docker" {
		if err := godotenv.Load("../.env.dev"); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	return nil
}
