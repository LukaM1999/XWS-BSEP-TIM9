package main

import (
	"dislinkt/comment_service/startup"
	cfg "dislinkt/comment_service/startup/config"
	"dislinkt/common/loggers"
	"os"
	"os/signal"
	"syscall"
)

var log = loggers.NewCommentLogger()

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		log.Info("Comment service stopped")
		done <- true
		os.Exit(0)
	}()
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	log.Info("Comment service started")
	server.Start()
	<-done
}
