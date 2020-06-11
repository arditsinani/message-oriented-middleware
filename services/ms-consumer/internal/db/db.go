package db

import (
	"context"
	"log"
	"mom/services/ms-consumer/config"
	"mom/services/ms-consumer/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Config *config.Config
	Client *mongo.Client
}

func (db *DB) getDB() *mongo.Database {
	return db.Client.Database(db.Config.Mongo.DatabaseName)
}

func (db *DB) getCollection(coll string) *mongo.Collection {
	return db.getDB().Collection(coll)
}

func (db *DB) Create(ctx context.Context, test interface{}, coll string) (interface{}, error) {
	_, err := db.getCollection(coll).InsertOne(ctx, test)
	if err != nil {
		return models.CreateTestForm{}, err
	}
	//test.ID = insertResult.InsertedID
	return test, nil
}

func (db *DB) Get(ctx context.Context, filter bson.M, coll string) ([]*models.Test, error) {
	findOptions := options.Find()
	findOptions.SetLimit(1000)
	cur, err := db.getCollection(coll).Find(ctx, filter, findOptions)
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

func (db *DB) GetCursor(ctx context.Context, filter bson.M, coll string, findOptions *options.FindOptions) (*mongo.Cursor, error) {
	return db.getCollection(coll).Find(ctx, filter, findOptions)
}

func (db *DB) GetById(ctx context.Context, id primitive.ObjectID, coll string) (interface{}, error) {
	test := models.Test{}
	err := db.getCollection(coll).FindOne(ctx, bson.M{"_id": id}).Decode(&test)
	if err != nil {
		return test, err
	}
	return test, nil
}

func (db *DB) Update(ctx context.Context, id primitive.ObjectID, test interface{}, coll string) (interface{}, error) {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	opts := options.Update().SetUpsert(true)
	_, err := db.getCollection(coll).UpdateOne(ctx, filter, bson.M{"$set": test}, opts)
	if err != nil {
		return test, err
	}
	return test, nil
}

func (db *DB) Delete(ctx context.Context, id primitive.ObjectID, coll string) (interface{}, error) {
	_, err := db.getCollection(coll).DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return models.Test{}, err
	}
	return nil, nil
}

func (db *DB) Stream(ctx context.Context, coll string, pipeline mongo.Pipeline, options *options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	return db.getCollection(coll).Watch(ctx, pipeline, options)
}

func (db *DB) GetStreamOptions() *options.ChangeStreamOptions {
	return options.ChangeStream().SetFullDocument(options.UpdateLookup)

}

func New(conf *config.Config) (*DB, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(conf.Mongo.Url()))
	ctx := context.TODO()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	db := &DB{
		Config: conf,
		Client: client,
	}
	return db, nil
}
