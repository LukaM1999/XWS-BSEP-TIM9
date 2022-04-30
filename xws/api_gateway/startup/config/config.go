package config

type Config struct {
	Port         string
	SecurityHost string
	SecurityPort string
}

func NewConfig() *Config {
	return &Config{
		Port:         "8000",
		SecurityHost: "security_service",
		SecurityPort: "8001",
	}
}
