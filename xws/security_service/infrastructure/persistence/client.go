package persistence

import (
	"context"
	"dislinkt/common/loggers"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var log = loggers.NewSecurityLogger()

func GetClient(host, port string) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s/", host, port)
	log.Info("Connecting to security database: ", uri)
	options := options.Client().ApplyURI(uri)
	return mongo.Connect(context.TODO(), options)
}
