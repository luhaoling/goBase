package topic

import (
	"fmt"
	"github.com/segmentio/kafka-go"
)

func Topic() {
	// 连接 kafka 列表
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	// 获取所有分区

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}
	// 遍历所有分区取 topic
	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	// 打印 topic
	for k := range m {
		fmt.Println(k)
	}
}
