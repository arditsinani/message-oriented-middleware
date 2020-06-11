package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mom/services/ms-consumer/internal/db"
	"mom/services/ms-consumer/internal/models"
)

type TestService struct {
	DB *db.DB
}

func (s *TestService) Create(ctx context.Context, test models.CreateTestForm, coll string) (models.Test, error) {
	result, err := s.DB.Create(ctx, test, coll)
	if err != nil {
		return models.Test{}, err
	}
	response := models.Test{}
	id := result.InsertedID.(primitive.ObjectID)
	response.ID = id
	response.Name = test.Name
	response.Surname = test.Surname
	return response, nil
}

func (s *TestService) Get(ctx context.Context, filter bson.M, coll string) ([]*models.Test, error) {
	findOptions := options.Find()
	findOptions.SetLimit(1000)
	cur, err := s.DB.GetCursor(ctx, filter, coll, findOptions)
	var results []*models.Test
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.Test
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results, err
}

func (s *TestService) GetById(ctx context.Context, id primitive.ObjectID, coll string) (models.Test, error) {
	test := models.Test{}
	singleResult := s.DB.GetById(ctx, id, coll)
	err := singleResult.Decode(&test)
	return test, err
}

func (s *TestService) Update(ctx context.Context, id primitive.ObjectID, test models.UpdateTestForm, coll string) (*mongo.UpdateResult, error){
	return s.DB.Update(ctx, id, test, coll)
}

func (s *TestService) Delete(ctx context.Context, id primitive.ObjectID, coll string) (*mongo.DeleteResult, error) {
	return s.DB.Delete(ctx, id, coll)
}

