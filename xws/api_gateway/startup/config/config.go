package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port           string
	SecurityHost   string
	SecurityPort   string
	ProfileHost    string
	ProfilePort    string
	CommentHost    string
	CommentPort    string
	ReactionHost   string
	ReactionPort   string
	ConnectionHost string
	ConnectionPort string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:           os.Getenv("GATEWAY_PORT"),
		SecurityHost:   os.Getenv("SECURITY_SERVICE_HOST"),
		SecurityPort:   os.Getenv("SECURITY_SERVICE_PORT"),
		ProfileHost:    os.Getenv("PROFILE_SERVICE_HOST"),
		ProfilePort:    os.Getenv("PROFILE_SERVICE_PORT"),
		CommentHost:    os.Getenv("COMMENT_SERVICE_HOST"),
		CommentPort:    os.Getenv("COMMENT_SERVICE_PORT"),
		ReactionHost:   os.Getenv("REACTION_SERVICE_HOST"),
		ReactionPort:   os.Getenv("REACTION_SERVICE_PORT"),
		ConnectionHost: os.Getenv("CONNECTION_SERVICE_HOST"),
		ConnectionPort: os.Getenv("CONNECTION_SERVICE_PORT"),
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
