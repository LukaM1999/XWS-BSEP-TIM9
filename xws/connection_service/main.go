package main

import (
	"dislinkt/common/loggers"
	"dislinkt/connection_service/startup"
	cfg "dislinkt/connection_service/startup/config"
	"os"
	"os/signal"
	"syscall"
)

var log = loggers.NewConnectionLogger()

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		log.Info("Connection service stopped")
		done <- true
		os.Exit(0)
	}()
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	log.Info("Connection service started")
	server.Start()
	<-done
}
