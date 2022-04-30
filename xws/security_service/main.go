package main

import (
	"dislinkt/security_service/startup"
	cfg "dislinkt/security_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
