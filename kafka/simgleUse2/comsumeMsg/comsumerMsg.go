package comsumeMsg

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func ReadByConn() {
	topic := "topicTest"
	partition := 0

	// 建立连接
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	// 设置读取超时时间
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	// 读取消息
	batch := conn.ReadBatch(10e3, 1e6)
	b := make([]byte, 10e3)
	for {
		fmt.Println("hello")
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		// 打印消息
		fmt.Println(string(b[:n]))
	}

	// 关闭读取连接
	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	// 关闭 kafka 连接
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}
