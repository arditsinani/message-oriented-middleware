package watchers

import (
	"context"
	"fmt"
	"log"
	"mom/services/ms-extractor/config"
	"mom/services/ms-extractor/internal/db"
	"mom/services/ms-extractor/internal/models"
	"mom/services/ms-extractor/pkg/kafka"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TestWatcher struct {
	Config *config.Config
	DB     *db.DB
	Kafka  *kafka.Kafka
}

func (w *TestWatcher) CreateStream() {
	// specify a pipeline that will only match "insert" events
	// specify the MaxAwaitTimeOption to have each attempt wait two seconds for new documents
	matchStage := bson.D{{"$match", bson.D{{"operationType", "insert"}}}}
	streamOptions := w.DB.GetStreamOptions()
	changeStream, err := w.DB.Stream(context.TODO(), models.TESTCOLLECTION, mongo.Pipeline{matchStage}, streamOptions)
	if err != nil {
		log.Fatal(err)
	}

	// print out all change stream events in the order they're received
	// see the mongo.ChangeStream documentation for more examples of using change streams
	for changeStream.Next(context.TODO()) {
		//w.Kafka.Producers.TestProducer.ProducerBatch(changeStream.Current, "test")
		w.Kafka.Producers.TestProducer.Producer(changeStream.Current, "test")
		//fmt.Println(changeStream.Current)
	}
}

func (w *TestWatcher) CreateTestStreamFromFind() {
	// Pass these options to the Find method
	findOptions := options.Find()

	// Here's an array in which you can store the decoded documents
	var results []*models.Test

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := w.DB.GetCursor(context.TODO(), bson.M{}, models.TESTCOLLECTION, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.Test
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(&elem)
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	//cur.Close(context.TODO())

	//fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
}
