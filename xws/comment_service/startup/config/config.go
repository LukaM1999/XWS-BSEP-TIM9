package config

import (
	"dislinkt/common/loggers"
	"github.com/joho/godotenv"
	"os"
)

var log = loggers.NewCommentLogger()

type Config struct {
	Port                        string
	CommentDBHost               string
	CommentDBPort               string
	NatsHost                    string
	NatsPort                    string
	NatsUser                    string
	NatsPass                    string
	UpdateProfileCommandSubject string
	UpdateProfileReplySubject   string
	DeletePostCommandSubject    string
	DeletePostReplySubject      string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:                        os.Getenv("COMMENT_SERVICE_PORT"),
		CommentDBHost:               os.Getenv("COMMENT_DB_HOST"),
		CommentDBPort:               os.Getenv("COMMENT_DB_PORT"),
		NatsHost:                    os.Getenv("NATS_HOST"),
		NatsPort:                    os.Getenv("NATS_PORT"),
		NatsUser:                    os.Getenv("NATS_USER"),
		NatsPass:                    os.Getenv("NATS_PASS"),
		UpdateProfileCommandSubject: os.Getenv("UPDATE_PROFILE_COMMAND_SUBJECT"),
		UpdateProfileReplySubject:   os.Getenv("UPDATE_PROFILE_REPLY_SUBJECT"),
		DeletePostCommandSubject:    os.Getenv("DELETE_POST_COMMAND_SUBJECT"),
		DeletePostReplySubject:      os.Getenv("DELETE_POST_REPLY_SUBJECT"),
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
