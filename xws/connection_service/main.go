package main

import (
	"dislinkt/connection_service/startup"
	cfg "dislinkt/connection_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
