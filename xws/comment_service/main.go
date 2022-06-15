package main

import (
	"dislinkt/comment_service/startup"
	cfg "dislinkt/comment_service/startup/config"
	"dislinkt/common/loggers"
)

var log = loggers.NewCommentLogger()

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	log.Info("Comment service started")
	server.Start()
}
