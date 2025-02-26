package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		GroupID:  "consumer-group-id",
		Topic:    "topic-A",
		MaxBytes: 10e6,
	})
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v:%s=%s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
