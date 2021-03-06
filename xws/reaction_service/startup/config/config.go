package config

import (
	"dislinkt/common/loggers"
	"github.com/joho/godotenv"
	"os"
)

var log = loggers.NewReactionLogger()

type Config struct {
	Port                     string
	ReactionDBHost           string
	ReactionDBPort           string
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	DeletePostCommandSubject string
	DeletePostReplySubject   string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:                     os.Getenv("REACTION_SERVICE_PORT"),
		ReactionDBHost:           os.Getenv("REACTION_DB_HOST"),
		ReactionDBPort:           os.Getenv("REACTION_DB_PORT"),
		NatsHost:                 os.Getenv("NATS_HOST"),
		NatsPort:                 os.Getenv("NATS_PORT"),
		NatsUser:                 os.Getenv("NATS_USER"),
		NatsPass:                 os.Getenv("NATS_PASS"),
		DeletePostCommandSubject: os.Getenv("DELETE_POST_COMMAND_SUBJECT"),
		DeletePostReplySubject:   os.Getenv("DELETE_POST_REPLY_SUBJECT"),
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
