package topic

import (
	"github.com/segmentio/kafka-go"
)

func CreateTopic() {
	// 连接 kafka 列表
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	// 获取所有分区

	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             "testTopic",
		NumPartitions:     1,
		ReplicationFactor: 1,
	})
	if err != nil {
		panic(err.Error())
	}

}
