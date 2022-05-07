package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port             string
	ConnectionDBHost string
	ConnectionDBPort string
	PostHost         string
	PostPort         string
	NatsHost         string
	NatsPort         string
	NatsUser         string
	NatsPass         string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:             os.Getenv("CONNECTION_SERVICE_PORT"),
		ConnectionDBHost: os.Getenv("CONNECTION_DB_HOST"),
		ConnectionDBPort: os.Getenv("CONNECTION_DB_PORT"),
		PostHost:         os.Getenv("POST_SERVICE_HOST"),
		PostPort:         os.Getenv("POST_SERVICE_PORT"),
		NatsHost:         os.Getenv("NATS_HOST"),
		NatsPort:         os.Getenv("NATS_PORT"),
		NatsUser:         os.Getenv("NATS_USER"),
		NatsPass:         os.Getenv("NATS_PASS"),
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
