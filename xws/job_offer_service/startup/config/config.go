package config

import (
	"dislinkt/common/loggers"
	"github.com/joho/godotenv"
	"os"
)

var log = loggers.NewConnectionLogger()

type Config struct {
	Port           string
	JobOfferDBHost string
	JobOfferDBPort string
	ProfileHost    string
	ProfilePort    string
	NatsHost       string
	NatsPort       string
	NatsUser       string
	NatsPass       string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:           os.Getenv("JOB_OFFER_SERVICE_PORT"),
		JobOfferDBHost: os.Getenv("JOB_OFFER_DB_HOST"),
		JobOfferDBPort: os.Getenv("JOB_OFFER_DB_PORT"),
		ProfileHost:    os.Getenv("PROFILE_SERVICE_HOST"),
		ProfilePort:    os.Getenv("PROFILE_SERVICE_PORT"),
		NatsHost:       os.Getenv("NATS_HOST"),
		NatsPort:       os.Getenv("NATS_PORT"),
		NatsUser:       os.Getenv("NATS_USER"),
		NatsPass:       os.Getenv("NATS_PASS"),
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
