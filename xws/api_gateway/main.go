package main

import (
	"dislinkt/api_gateway/startup"
	cfg "dislinkt/api_gateway/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
