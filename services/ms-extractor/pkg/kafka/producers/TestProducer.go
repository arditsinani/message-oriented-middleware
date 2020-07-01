package kafka

import (
	"context"
	"fmt"
	"mom/services/ms-extractor/config"
	"mom/services/ms-extractor/internal/db"
	"time"

	"github.com/segmentio/kafka-go"
)

type TestProducer struct {
	Config *config.Config
	DB  *db.DB
}

func (p *TestProducer) Producer(raw db.Raw, topic string) {
	fmt.Println("started producer")
	// to produce messages
	//topic := "test"
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), p.Config.Kafka.Network, p.Config.Kafka.Address(), topic, partition)

	bytes, err := conn.Write(raw)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("writen bytes => ", bytes)

	conn.Close()
}
func (p *TestProducer) ProducerBatch(raw db.Raw, topic string) {
	fmt.Println("started producer")
	// to produce messages
	//topic := "test"
	//partition := 0

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{p.Config.Kafka.Address()},
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 10 * time.Millisecond,
	})

	defer writer.Close()
	err := writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: raw,
		})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("writen message => ", raw)

}

//func EventListenerTestConfluent() {
//
//	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
//	if err != nil {
//		panic(err)
//	}
//
//	defer p.Close()
//
//	// Delivery report handler for produced messages
//	go func() {
//		for e := range p.Events() {
//			switch ev := e.(type) {
//			case *kafka.Message:
//				if ev.TopicPartition.Error != nil {
//					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
//				} else {
//					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
//				}
//			}
//		}
//	}()
//	// Wait for message deliveries before shutting down
//	p.Flush(15 * 1000)
//}

//func ProduceTestMessagesConfluent() {
//	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
//	if err != nil {
//		panic(err)
//	}
//
//	defer p.Close()
//
//	// Produce messages to topic (asynchronously)
//	topic := "test"
//	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
//		err = p.Produce(&kafka.Message{
//			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
//			Value:          []byte(word),
//		}, nil)
//		if err != nil {
//			panic(err)
//		}
//		fmt.Printf("produced %v\n", word)
//	}
//
//	// Wait for message deliveries before shutting down
//	p.Flush(15 * 1000)
//
//}
