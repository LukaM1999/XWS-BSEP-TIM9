package main

import (
	"dislinkt/common/loggers"
	"dislinkt/reaction_service/startup"
	cfg "dislinkt/reaction_service/startup/config"
)

var log = loggers.NewReactionLogger()

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	log.Info("Reaction service started")
	server.Start()
}
