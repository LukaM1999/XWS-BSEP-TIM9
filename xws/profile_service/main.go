package main

import (
	"dislinkt/profile_service/startup"
	cfg "dislinkt/profile_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
