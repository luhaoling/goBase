package producer

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
	Producer *kafka.Writer
)

func init() {
	Producer = &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  topic,
		Balancer:               &kafka.Hash{},
		MaxAttempts:            0,
		WriteBackoffMin:        0,
		WriteBackoffMax:        0,
		BatchSize:              0,
		BatchBytes:             0,
		BatchTimeout:           0,
		ReadTimeout:            0,
		WriteTimeout:           time.Second,
		RequiredAcks:           kafka.RequireNone,
		Async:                  false,
		Completion:             nil,
		Compression:            0,
		Logger:                 nil,
		ErrorLogger:            nil,
		Transport:              nil,
		AllowAutoTopicCreation: false,
	}
}

func SendMessage(ctx context.Context, user *model.User) {
	msgContent, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("json unmarshal user err,user:%v,err:%v", user, err)
	}
	msg := kafka.Message{
		Topic:         "",
		Partition:     0,
		Offset:        0,
		HighWaterMark: 0,
		Key:           []byte(fmt.Sprintf("%d", user.Id)),
		Value:         msgContent,
		Headers:       nil,
		WriterData:    nil,
		Time:          time.Time{},
	}
	err = Producer.WriteMessages(ctx, msg)
	if err != nil {
		fmt.Printf("写入 kafka 失败，user:%v,err:%v", user, err)
	}

}
