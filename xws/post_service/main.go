package main

import (
	"dislinkt/post_service/startup"
	cfg "dislinkt/post_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
