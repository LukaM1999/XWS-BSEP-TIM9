package config

import "os"

type Config struct {
	Port          string
	CommentDBHost string
	CommentDBPort string
	NatsHost      string
	NatsPort      string
	NatsUser      string
	NatsPass      string
}

func NewConfig() *Config {
	return &Config{
		Port:          "8001",
		CommentDBHost: "comment_db",
		CommentDBPort: "27017",
		NatsHost:      os.Getenv("NATS_HOST"),
		NatsPort:      os.Getenv("NATS_PORT"),
		NatsUser:      os.Getenv("NATS_USER"),
		NatsPass:      os.Getenv("NATS_PASS"),
	}
}
