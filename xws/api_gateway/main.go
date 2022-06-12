package main

import (
	"dislinkt/api_gateway/startup"
	cfg "dislinkt/api_gateway/startup/config"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"os/signal"
	"syscall"
)

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
	initLogger()
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	log.Info("Security service started")
	server.Start()
	<-done
}

func initLogger() {
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "../../logs/api_gateway/api_gateway.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
	log.SetOutput(multiWriter)
}
