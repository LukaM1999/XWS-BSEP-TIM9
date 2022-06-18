package main

import (
	"dislinkt/api_gateway/startup"
	cfg "dislinkt/api_gateway/startup/config"
	"dislinkt/common/loggers"
	"os"
	"os/signal"
	"syscall"
)

var log = loggers.NewGatewayLogger()

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		log.Info("API Gateway stopped")
		done <- true
		os.Exit(0)
	}()
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	log.Info("API Gateway started")
	server.Start()
	<-done
}
