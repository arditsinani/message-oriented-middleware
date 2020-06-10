package main

import (
	"fmt"
	"mom/services/ms-consumer/config"
	"mom/services/ms-consumer/internal/db"
	"mom/services/ms-consumer/pkg/kafka"
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
	_ = kafka.New(conf, mongo)
}
