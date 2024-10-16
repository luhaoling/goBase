package sendMsg

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func WriteByConn() {
	topic := "topicTest"
	parrtition := 0

	// 连接
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, parrtition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	// 设置最迟发送时间
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	// 往 kafka 中发送信息
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write message:", err)
	}

	// 关闭连接
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	fmt.Println("sendSuccess")
}
