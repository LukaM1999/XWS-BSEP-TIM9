package config

import (
	"dislinkt/common/loggers"
	"github.com/joho/godotenv"
	"os"
)

var log = loggers.NewConnectionLogger()

type Config struct {
	Port                     string
	JobOfferDBHost           string
	JobOfferDBPort           string
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	PromoteJobCommandSubject string
	PromoteJobReplySubject   string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:                     os.Getenv("JOB_OFFER_SERVICE_PORT"),
		JobOfferDBHost:           os.Getenv("JOB_OFFER_DB_HOST"),
		JobOfferDBPort:           os.Getenv("JOB_OFFER_DB_PORT"),
		NatsHost:                 os.Getenv("NATS_HOST"),
		NatsPort:                 os.Getenv("NATS_PORT"),
		NatsUser:                 os.Getenv("NATS_USER"),
		NatsPass:                 os.Getenv("NATS_PASS"),
		PromoteJobCommandSubject: os.Getenv("PROMOTE_JOB_COMMAND_SUBJECT"),
		PromoteJobReplySubject:   os.Getenv("PROMOTE_JOB_REPLY_SUBJECT"),
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
