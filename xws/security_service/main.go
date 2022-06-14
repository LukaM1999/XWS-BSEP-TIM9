package main

import (
	"dislinkt/common/loggers"
	"dislinkt/security_service/startup"
	cfg "dislinkt/security_service/startup/config"
	"os"
	"os/signal"
	"syscall"
)

var log = loggers.NewSecurityLogger()

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		log.Info("Security service stopped")
		done <- true
		os.Exit(0)
	}()
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	log.Info("Security service started")
	server.Start()
	<-done
}
