package kafka

import (
	"mom/services/ms-extractor/config"
	"mom/services/ms-extractor/internal/db"
	producers "mom/services/ms-extractor/pkg/kafka/producers"
)

type Kafka struct {
	Producers Producers
}

type Producers struct {
	TestP producers.TestP
}

func New(conf *config.Config, db *db.DB) *Kafka {
	kafka := Kafka{
		Producers: Producers{producers.TestP{Config: conf, DB: db}},
	}

	return &kafka
}
