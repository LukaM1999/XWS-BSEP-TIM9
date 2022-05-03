package main

import (
	"dislinkt/reaction_service/startup"
	cfg "dislinkt/reaction_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
