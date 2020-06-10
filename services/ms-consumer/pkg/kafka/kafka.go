package kafka

import (
	"go.mongodb.org/mongo-driver/mongo"
	"mom/services/ms-consumer/config"
	consumers "mom/services/ms-consumer/pkg/kafka/consumers"
)

type Kafka struct {
	Consumers Consumers
}

type Consumers struct {
	TestConsumer consumers.TestConsumer
}

func New(conf config.Config, mongo *mongo.Client) *Kafka {
	kafka := Kafka{
		Consumers: Consumers{consumers.TestConsumer{Config: conf, Mongo: mongo}},
	}
	kafka.Consumers.TestConsumer.Consumer()
	//go kafka.Consumers.TestConsumer.ReadTopic()
	return &kafka
}