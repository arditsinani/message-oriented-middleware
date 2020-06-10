package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mom/services/ms-consumer/internal/models"
)

func Insert(c *mongo.Client, ctx context.Context, test interface{}, db string, collection string) (interface{}, error){
	_, err := c.Database(db).Collection(collection).InsertOne(ctx, test)
	if err != nil {
		return models.CreateTestForm{}, err
	}
	//test.ID = insertResult.InsertedID
	return test, nil
}

func Get(c *mongo.Client, ctx context.Context, filter bson.M, db string, collection string) ([]*models.Test, error) {
	findOptions := options.Find()
	findOptions.SetLimit(1000)
	cur, err := c.Database(db).Collection(collection).Find(ctx, filter, findOptions)
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

func GetById(c *mongo.Client, ctx context.Context, id primitive.ObjectID, db string, collection string)  (interface{}, error) {
	test := models.Test{}
	err := c.Database(db).Collection(collection).FindOne(ctx, bson.M{"_id": id}).Decode(&test)
	if err != nil {
		return test, err
	}
	return test, nil
}

func Update(c *mongo.Client, ctx context.Context, id primitive.ObjectID, test interface{}, db string, collection string) (interface{}, error) {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	opts := options.Update().SetUpsert(true)
	_, err := c.Database(db).Collection(collection).UpdateOne(ctx, filter, bson.M{ "$set": test }, opts)
	if err != nil {
		return test, err
	}
	return test, nil
}

func Delete(c *mongo.Client, ctx context.Context, id primitive.ObjectID, db string, collection string)  (interface{}, error) {
	_, err := c.Database(db).Collection(collection).DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return models.Test{}, err
	}
	return nil, nil
}

