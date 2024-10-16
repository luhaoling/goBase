package main

import (
	"github.com/hibiken/asynq"
	"goBase/asynq/single_example/test_delivery"
	"log"
)

func main() {
	// 连接 redis
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: "localhost:6379",
			DB:   3,
		},
		asynq.Config{
			Concurrency: 5,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	mux := asynq.NewServeMux()
	// 收到消息后，由下面这个函数处理
	mux.HandleFunc(test_delivery.TypeEmailDelivery, test_delivery.HandleEmailDeliveryTask)
	if err := srv.Run(mux); err != nil {
		log.Fatalf("cound not run server:%v", err)
	}
}
