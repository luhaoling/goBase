package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"goBase/kafka/simgleUse/model"
	"time"
)

var (
	topic    = "test"
	Consumer *kafka.Reader
)

func init() {
	Consumer = kafka.NewReader(kafka.ReaderConfig{
		Brokers:                []string{"localhost:9092"},
		GroupID:                "test",
		GroupTopics:            nil,
		Topic:                  topic,
		Partition:              0,
		Dialer:                 nil,
		QueueCapacity:          0,
		MinBytes:               0,
		MaxBytes:               0,
		MaxWait:                0,
		ReadBatchTimeout:       0,
		ReadLagInterval:        0,
		GroupBalancers:         nil,
		HeartbeatInterval:      0,
		CommitInterval:         time.Second,
		PartitionWatchInterval: 0,
		WatchPartitionChanges:  false,
		SessionTimeout:         0,
		RebalanceTimeout:       0,
		JoinGroupBackoff:       0,
		RetentionTime:          0,
		StartOffset:            kafka.FirstOffset,
		ReadBackoffMin:         0,
		ReadBackoffMax:         0,
		Logger:                 nil,
		ErrorLogger:            nil,
		IsolationLevel:         0,
		MaxAttempts:            0,
		OffsetOutOfRangeError:  false,
	})
}

func ReadMessage(ctx context.Context) {
	for {
		if msg, err := Consumer.ReadMessage(ctx); err != nil {
			fmt.Printf("读 kafka 失败，err:%v", err)
			continue
		} else {
			user := &model.User{}
			fmt.Println("！！！！", string(msg.Value))
			err := json.Unmarshal(msg.Value, user)
			if err != nil {
				fmt.Printf("json unmarshal msg value err,msg:%v,err:%v", user, err)
				continue
			}
			fmt.Printf("topic=%s,partition=%d,offset=%d,key=%s,user=%v\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, user)
		}
	}
}
