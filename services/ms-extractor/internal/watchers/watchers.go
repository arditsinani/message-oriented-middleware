package watchers

import (
	"mom/services/ms-extractor/config"
	"mom/services/ms-extractor/internal/db"
	"mom/services/ms-extractor/pkg/kafka"
	"reflect"
)

type Watchers struct {
	Test TestW
}

func New(conf *config.Config, db *db.DB, kafka *kafka.Kafka) {
	watchers := Watchers{Test: TestW{Config: conf, DB: db, Kafka: kafka}}

	values := reflect.ValueOf(&watchers).Elem()
	for i := 0; i < values.NumField(); i++ {
		watcher := values.Field(i).Addr()
		go watcher.MethodByName("CreateStream").Call([]reflect.Value{})
	}

}
