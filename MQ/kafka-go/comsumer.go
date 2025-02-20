package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	topic := "my-topic"
	partition := 0

	// 连接Kafka
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	defer conn.Close()

	// 设置超时（30秒）
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))

	// 调整批次参数
	batch := conn.ReadBatch(10e3, 100e3)
	defer func() {
		_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := batch.Close(); err != nil {
			log.Fatal("关闭批次失败:", err)
		}
	}()

	// 读取消息

	for {
		msg, err := batch.ReadMessage()
		if err != nil {
			log.Fatal("读取消息错误:", err)
			break
		}
		fmt.Println(string(msg.Value))
	}
}
