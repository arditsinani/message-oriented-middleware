package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"mom/services/ms-extractor/config"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
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
