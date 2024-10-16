package main

import (
	"context"
	"fmt"
	"goBase/kafka/simgleUse/model"
	"goBase/kafka/simgleUse/producer"
)

// 发送者发送消息

func main() {
	ctx := context.Background()
	for i := 0; i < 5; i++ {
		user := &model.User{
			Id:       int64(i + 1),
			UserName: fmt.Sprintf("lym:%d", i),
			Age:      18,
		}
		producer.SendMessage(ctx, user)
	}
	producer.Producer.Close()
}
