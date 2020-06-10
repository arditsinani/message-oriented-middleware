package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mom/services/ms-consumer/config"
)

func New(conf config.Config) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(conf.Mongo.Url()))
	ctx := context.TODO()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}
