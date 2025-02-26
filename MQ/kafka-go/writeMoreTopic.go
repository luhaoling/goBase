package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "topic-F",
		Balancer: &kafka.LeastBytes{},
	}
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Partition: 0,
			Key:       []byte("key-A0"),
			Value:     []byte("hello world"),
		},
		kafka.Message{
			Partition: 1,
			Key:       []byte("key-A1"),
			Value:     []byte("hello world"),
		},
		kafka.Message{
			Partition: 2,
			Key:       []byte("key-A2"),
			Value:     []byte("hello world"),
		},
	)
	if err != nil {
		log.Fatal("failed to writer message:", err)
	}
	fmt.Println("send success")
}
