package main

import (
	"fmt"
	"mom/services/ms-extractor/config"
	"mom/services/ms-extractor/internal/db"
	"mom/services/ms-extractor/internal/server"
	"mom/services/ms-extractor/internal/watchers"
	"mom/services/ms-extractor/pkg/kafka"
)

func main() {
	// init config
	conf := config.New()
	// init database
	mongo, err := db.New(conf)
	if err != nil {
		fmt.Print(err)
	}
	// init kafka
	kafka := kafka.New(conf, mongo)
	// init watchers
	watchers.New(conf, mongo, kafka)
	// init server
	server.New(conf, mongo)
}
