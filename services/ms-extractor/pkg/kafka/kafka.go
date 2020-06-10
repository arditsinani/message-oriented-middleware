package kafka

import (
	"go.mongodb.org/mongo-driver/mongo"
	"mom/services/ms-extractor/config"
	producers "mom/services/ms-extractor/pkg/kafka/producers"
)

type Kafka struct {
	Producers Producers
}

type Producers struct {
	TestProducer producers.TestProducer
}

func New(conf config.Config, mongo *mongo.Client) *Kafka {
	kafka := Kafka{
		Producers: Producers{producers.TestProducer{Config: conf, Mongo: mongo}},
	}

	return &kafka
}