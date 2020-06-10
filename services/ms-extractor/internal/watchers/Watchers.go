package watchers

import (
	"go.mongodb.org/mongo-driver/mongo"
	"mom/services/ms-extractor/config"
	"mom/services/ms-extractor/pkg/kafka"
	"reflect"
)

type Watchers struct {
	Test TestWatcher
}

func New(conf config.Config,mongo *mongo.Client,kafka *kafka.Kafka ) {
	watchers := Watchers{Test: TestWatcher{Config: conf, Mongo: mongo, Kafka: kafka}}

	values := reflect.ValueOf(&watchers).Elem()
	for i := 0; i < values.NumField(); i++ {
		watcher := values.Field(i).Addr()
		go watcher.MethodByName("CreateStream").Call([]reflect.Value{})
	}

}

