package main

import (
	"dislinkt/common/loggers"
	"dislinkt/connection_service/startup"
	cfg "dislinkt/connection_service/startup/config"
)

var log = loggers.NewConnectionLogger()

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	log.Info("Connection service started")
	server.Start()
}
