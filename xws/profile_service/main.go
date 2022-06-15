package main

import (
	"dislinkt/common/loggers"
	"dislinkt/profile_service/startup"
	cfg "dislinkt/profile_service/startup/config"
)

var log = loggers.NewProfileLogger()

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	log.Info("Profile service started")
	server.Start()
}
