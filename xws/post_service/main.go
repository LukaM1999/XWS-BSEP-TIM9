package main

import (
	"dislinkt/common/loggers"
	"dislinkt/post_service/startup"
	cfg "dislinkt/post_service/startup/config"
)

var log = loggers.NewPostLogger()

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	log.Info("Post service started")
	server.Start()
}
