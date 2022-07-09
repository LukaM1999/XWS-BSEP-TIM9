package loggers

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var securityLogger = logrus.New()
var gatewayLogger = logrus.New()
var interceptorLogger = logrus.New()
var profileLogger = logrus.New()
var postLogger = logrus.New()
var commentLogger = logrus.New()
var connectionLogger = logrus.New()
var reactionLogger = logrus.New()
var jobLogger = logrus.New()

func NewJobLogger() *logrus.Logger {
	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	jobLogger.SetLevel(logrus.InfoLevel)
	jobLogger.SetReportCaller(true)
	jobLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   logPathPrefix + "job_offer_service/job_offer.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
	jobLogger.SetOutput(multiWriter)
	return jobLogger
}

func NewReactionLogger() *logrus.Logger {
	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	reactionLogger.SetLevel(logrus.InfoLevel)
	reactionLogger.SetReportCaller(true)
	reactionLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   logPathPrefix + "reaction_service/reaction.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
	reactionLogger.SetOutput(multiWriter)
	return reactionLogger
}

func NewConnectionLogger() *logrus.Logger {
	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	connectionLogger.SetLevel(logrus.InfoLevel)
	connectionLogger.SetReportCaller(true)
	connectionLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   logPathPrefix + "connection_service/connection.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
	connectionLogger.SetOutput(multiWriter)
	return connectionLogger
}

func NewCommentLogger() *logrus.Logger {
	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	commentLogger.SetLevel(logrus.InfoLevel)
	commentLogger.SetReportCaller(true)
	commentLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   logPathPrefix + "comment_service/comment.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
	commentLogger.SetOutput(multiWriter)
	return commentLogger
}

func NewPostLogger() *logrus.Logger {
	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	postLogger.SetLevel(logrus.InfoLevel)
	postLogger.SetReportCaller(true)
	postLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   logPathPrefix + "post_service/post.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
	postLogger.SetOutput(multiWriter)
	return postLogger
}

func NewProfileLogger() *logrus.Logger {
	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	profileLogger.SetLevel(logrus.InfoLevel)
	profileLogger.SetReportCaller(true)
	profileLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   logPathPrefix + "profile_service/profile.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
	profileLogger.SetOutput(multiWriter)
	return profileLogger
}

func NewSecurityLogger() *logrus.Logger {
	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	securityLogger.SetLevel(logrus.InfoLevel)
	securityLogger.SetReportCaller(true)
	securityLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   logPathPrefix + "security_service/security.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
	securityLogger.SetOutput(multiWriter)
	return securityLogger
}

func NewGatewayLogger() *logrus.Logger {
	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	gatewayLogger.SetLevel(logrus.InfoLevel)
	gatewayLogger.SetReportCaller(true)
	gatewayLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   logPathPrefix + "api_gateway/api_gateway.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
	gatewayLogger.SetOutput(multiWriter)
	return gatewayLogger
}

func NewInterceptorLogger() *logrus.Logger {
	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	interceptorLogger.SetLevel(logrus.InfoLevel)
	interceptorLogger.SetReportCaller(true)
	interceptorLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   logPathPrefix + "auth_interceptor/auth_interceptor.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
	interceptorLogger.SetOutput(multiWriter)
	return interceptorLogger
}
