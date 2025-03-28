package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	queue "goBase/redis/releaseMQ"
	"log"
)

var RedisQueue *queue.Queue
var Redis *redis.Client

func main() {
	option := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
	client := redis.NewClient(option)

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(fmt.Sprintf("Connect to redis:%v", err))
	}
	Redis = client
	InitRedisQueue()
}

func InitRedisQueue() {
	RedisQueue = queue.NewQueue(context.Background(), Redis, queue.WithTopic("topic"))
	RedisQueue.Start()
}
