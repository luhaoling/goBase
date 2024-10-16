package client

import (
	"fmt"
	"github.com/hibiken/asynq"
	"goBase/asynq/single_example/test_delivery"
	"log"
	"time"
)

// 添加一个消息
func EmailDeliveryTaskAdd(i int) {
	// 连接 redis
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: "localhost:6379",
		DB:   3,
	})
	defer client.Close()
	// 创建一个任务
	task, err := test_delivery.NewEmailDeliveryTask(42, fmt.Sprintf("some:template:id:%d", i), `{"name":"lisi"}`)
	if err != nil {
		log.Fatalf("cound not create task:%v", err)
	}
	// 将任务加入到队列中，5s 后处理
	info, err := client.Enqueue(task, asynq.ProcessIn(5*time.Second))
	if err != nil {
		log.Fatalf("cound not enqueue task:%v", err)
	}
	log.Printf("enqueued task:id=%s queue=%s", info.ID, info.Queue)
}
