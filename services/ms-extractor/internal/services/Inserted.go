package services

import (
	"go.mongodb.org/mongo-driver/mongo"
	"mom/services/ms-extractor/internal/models"
	"context"
)

func Inserted(c *mongo.Client, ctx context.Context, test interface{}, db string, collection string) (interface{}, error){
	_, err := c.Database(db).Collection(collection).InsertOne(ctx, test)
	if err != nil {
		return models.CreateTestForm{}, err
	}
	return test, nil
}

