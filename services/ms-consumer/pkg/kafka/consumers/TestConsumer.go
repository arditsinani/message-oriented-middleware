package kafka

import (
	"context"
	"fmt"
	"log"
	"mom/services/ms-consumer/config"
	"mom/services/ms-consumer/internal/db"
	"mom/services/ms-consumer/internal/models"
	"time"

	"github.com/segmentio/kafka-go"
)

type TestConsumer struct {
	Config *config.Config
	DB     *db.DB
}

func (c *TestConsumer) Consumer() {
	fmt.Println("started consumer")
	topic := "test"

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{c.Config.Kafka.Address()},
		GroupID:  "mongo",
		Topic:    topic,
		MinBytes: 1,
		MaxBytes: 10e6, // 10MB
		MaxWait:  1 * time.Millisecond,
	})
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		_, err = c.DB.Create(context.Background(), interface{}(msg.Value), models.TESTCOPYCOLLECTION)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a single document: ", msg.Value)
	}
}

func (c *TestConsumer) ConsumerBatch() {
	fmt.Println("started consumer batch")
	topic := "test"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), c.Config.Kafka.Network, c.Config.Kafka.Address(), topic, partition)

	if err != nil {
		fmt.Println(err)
	}
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println("message =>" + string(b))
	}

	batch.Close()
	conn.Close()
}

//func (c *TestConsumer) ReadTopic() {
//	fmt.Println("started read topic")
//	topic := "test"
//	partition :=0
//	kafka
//	reader := kafka.NewReader(kafka.ReaderConfig{
//		Brokers:  []string{c.Config.Kafka.Address()},
//		GroupID:  "mongo",
//		Topic:    topic,
//		MinBytes: 1,
//		MaxBytes: 10e6, // 10MB
//		MaxWait: 1 * time.Millisecond,
//	})
//	defer reader.Close()
//	messages := reader.FetchMessage(context.Background())
//	for  {
//		msg := messages.Value
//	}
//}
