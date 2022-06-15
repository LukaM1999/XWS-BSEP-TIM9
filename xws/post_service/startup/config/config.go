package config

import (
	"dislinkt/common/loggers"
	"github.com/joho/godotenv"
	"os"
)

var log = loggers.NewPostLogger()

type Config struct {
	Port         string
	PostDBHost   string
	PostDBPort   string
	CommentHost  string
	CommentPort  string
	ReactionHost string
	ReactionPort string
	NatsHost     string
	NatsPort     string
	NatsUser     string
	NatsPass     string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:         os.Getenv("POST_SERVICE_PORT"),
		PostDBHost:   os.Getenv("POST_DB_HOST"),
		PostDBPort:   os.Getenv("POST_DB_PORT"),
		CommentHost:  os.Getenv("COMMENT_SERVICE_HOST"),
		CommentPort:  os.Getenv("COMMENT_SERVICE_PORT"),
		ReactionHost: os.Getenv("REACTION_SERVICE_HOST"),
		ReactionPort: os.Getenv("REACTION_SERVICE_PORT"),
		NatsHost:     os.Getenv("NATS_HOST"),
		NatsPort:     os.Getenv("NATS_PORT"),
		NatsUser:     os.Getenv("NATS_USER"),
		NatsPass:     os.Getenv("NATS_PASS"),
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
