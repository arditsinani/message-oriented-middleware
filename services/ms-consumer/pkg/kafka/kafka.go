package kafka

import (
	"mom/services/ms-consumer/config"
	"mom/services/ms-consumer/internal/db"
	consumers "mom/services/ms-consumer/pkg/kafka/consumers"
)

type Kafka struct {
	Consumers Consumers
}

type Consumers struct {
	TestConsumer consumers.TestConsumer
}

func New(conf *config.Config, db *db.DB) *Kafka {
	kafka := Kafka{
		Consumers: Consumers{consumers.TestConsumer{Config: conf, DB: db}},
	}
	kafka.Consumers.TestConsumer.Consumer()
	//go kafka.Consumers.TestConsumer.ReadTopic()
	return &kafka
}
