package config

import (
	"dislinkt/common/loggers"
	"github.com/joho/godotenv"
	"os"
)

var log = loggers.NewSecurityLogger()

type Config struct {
	Port                      string
	SecurityDBHost            string
	SecurityDBPort            string
	ProfileHost               string
	ProfilePort               string
	NatsHost                  string
	NatsPort                  string
	NatsUser                  string
	NatsPass                  string
	CreateOrderCommandSubject string
	CreateOrderReplySubject   string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:                      os.Getenv("SECURITY_SERVICE_PORT"),
		SecurityDBHost:            os.Getenv("SECURITY_DB_HOST"),
		SecurityDBPort:            os.Getenv("SECURITY_DB_PORT"),
		ProfileHost:               os.Getenv("PROFILE_SERVICE_HOST"),
		ProfilePort:               os.Getenv("PROFILE_SERVICE_PORT"),
		NatsHost:                  os.Getenv("NATS_HOST"),
		NatsPort:                  os.Getenv("NATS_PORT"),
		NatsUser:                  os.Getenv("NATS_USER"),
		NatsPass:                  os.Getenv("NATS_PASS"),
		CreateOrderCommandSubject: os.Getenv("CREATE_ORDER_COMMAND_SUBJECT"),
		CreateOrderReplySubject:   os.Getenv("CREATE_ORDER_REPLY_SUBJECT"),
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
