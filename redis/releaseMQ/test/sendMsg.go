package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
	queue "goBase/redis/releaseMQ"
	"log"
	"time"
)

func main() {
	id := uuid.NewV4().String()
	logMsg := queue.NewMessage(id, time.Now(), map[string]interface{}{"user_id": 123, "action": "login"})

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

	qu := queue.NewQueue(context.Background(), client)
	_, err = qu.Publish(logMsg)
	if err != nil {
		fmt.Println("send ", err)
	}
}
