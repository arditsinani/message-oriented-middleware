package watchers

import (
	"context"
	"log"
	"mom/services/ms-extractor/config"
	"mom/services/ms-extractor/internal/db"
	"mom/services/ms-extractor/internal/models"
	"mom/services/ms-extractor/pkg/kafka"
)

type TestW struct {
	Config *config.Config
	DB     *db.DB
	Kafka  *kafka.Kafka
}

func (w *TestW) CreateStream() {
	// specify a pipeline that will only match "insert" events
	// specify the MaxAwaitTimeOption to have each attempt wait two seconds for new documents
	matchStage := db.DType{{"$match", db.DType{{"operationType", "insert"}}}}
	streamOptions := w.DB.GetStreamOptions()
	changeStream, err := w.DB.Stream(context.TODO(), models.TESTCOLLECTION, db.Pipeline{matchStage}, streamOptions)
	if err != nil {
		log.Fatal(err)
	}

	// print out all change stream events in the order they're received
	// see the mongo.ChangeStream documentation for more examples of using change streams
	for changeStream.Next(context.TODO()) {
		//w.Kafka.Producers.TestProducer.ProducerBatch(changeStream.Current, "test")
		w.Kafka.Producers.TestP.Producer(changeStream.Current, "test")
		//fmt.Println(changeStream.Current)
	}
}
