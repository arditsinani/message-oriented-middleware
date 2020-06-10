package config

import (
	"fmt"
	"os"
)

type Config struct {
	Mongo 	Mongo
	Kafka 	Kafka
	Server 	Server
}

type Mongo struct {
	Hostname     string `default:"localhost" envconfig:"MONGODB_HOSTNAME"`
	Port         string `default:"30001" envconfig:"MONGODB_PORT"`
	DatabaseName string `default:"consumer" envconfig:"CONSUMER_DATABASE"`
	ReplicaSet   string `default:"rs1" envconfig:"MONGODB_REPLICASET"`
}

type Kafka struct {
	Hostname string `default:"kafka1" envconfig:"KAFKA_HOSTNAME"`
	Port     string `default:"9092" envconfig:"KAFKA_PORT"`
	Network  string `default:"tcp" envconfig:"KAFKA_NETWORK"`
}

type Server struct {
	Port string `default:":8080" envconfig:"SERVER_PORT"`
}

func (k *Kafka) Address() string {
	if os.Getenv("KAFKA_URI") != "" {
		return os.Getenv("KAFKA_URI")
	}
	return fmt.Sprintf("%s:%s", k.Hostname, k.Port)
}

//const MONGODB_URI ="mongodb://localhost:30001,localhost:30002,localhost:30003/?replicaSet=rs1&connect=direct"
//const MONGODB_URI ="mongodb://localhost:30001/?replicaSet=rs1&connect=direct"
func (m *Mongo) Url() string {
	if os.Getenv("MONGO_URI") != "" {
		return os.Getenv("MONGO_URI")
	}
	return fmt.Sprintf("mongodb://%s:%s/?replicaSet=%s&connect=direct", m.Hostname, m.Port, m.ReplicaSet)
}

func New() Config {
	return Config {
		Mongo: 	Mongo{Hostname: "localhost",Port: "30001",DatabaseName: "consumer",ReplicaSet: "rs1"},
		Kafka: 	Kafka{Hostname: "kafka1",Port: "9092", Network: "tcp"},
		Server: Server{Port: ":8080"},
	}
}