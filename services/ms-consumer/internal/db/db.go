package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mom/services/ms-consumer/config"
)

type DB struct {
	Config *config.Config
	Client *mongo.Client
}

// inherited types from mongo
type MType = bson.M
type AType = bson.A
type DType = bson.D
type EType = bson.E
type Raw = bson.Raw
type Pipeline = mongo.Pipeline
type ObjectID = primitive.ObjectID
type UpdateResult = mongo.UpdateResult
type DeleteResult = mongo.DeleteResult


func (db *DB) getDB() *mongo.Database {
	return db.Client.Database(db.Config.Mongo.DatabaseName)
}

func (db *DB) getCollection(coll string) *mongo.Collection {
	return db.getDB().Collection(coll)
}

func (db *DB) Create(ctx context.Context, doc interface{}, coll string) (*mongo.InsertOneResult, error) {
	return db.getCollection(coll).InsertOne(ctx, doc)
}

func (db *DB) GetCursor(ctx context.Context, filter bson.M, coll string, findOptions *options.FindOptions) (*mongo.Cursor, error) {
	return db.getCollection(coll).Find(ctx, filter, findOptions)
}

func (db *DB) GetById(ctx context.Context, id primitive.ObjectID, coll string) *mongo.SingleResult {
	return db.getCollection(coll).FindOne(ctx, bson.M{"_id": id})
}

func (db *DB) Update(ctx context.Context, id primitive.ObjectID, doc interface{}, coll string) (*UpdateResult, error) {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	opts := options.Update().SetUpsert(true)
	return db.getCollection(coll).UpdateOne(ctx, filter, bson.M{"$set": doc}, opts)
}

func (db *DB) Delete(ctx context.Context, id primitive.ObjectID, coll string) (*DeleteResult, error)  {
	return db.getCollection(coll).DeleteOne(ctx, bson.M{"_id": id})
}

func (db *DB) Stream(ctx context.Context, coll string, pipeline mongo.Pipeline, options *options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	return db.getCollection(coll).Watch(ctx, pipeline, options)
}

func (db *DB) GetStreamOptions() *options.ChangeStreamOptions {
	return options.ChangeStream().SetFullDocument(options.UpdateLookup)

}

func (db *DB) GetFindOptions() *options.FindOptions {
	return options.Find()
}

func (db *DB) GetObjectIDFromHex(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
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
