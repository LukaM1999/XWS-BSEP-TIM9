package main

import (
	"dislinkt/comment_service/startup"
	cfg "dislinkt/comment_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
