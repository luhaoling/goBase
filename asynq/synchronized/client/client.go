package main

import (
	"fmt"
	"github.com/hibiken/asynq"
	"goBase/asynq/synchronized/tasks"
	"log"
	"time"
)

type User struct {
	UserID int
}

func main() {
	r := asynq.RedisClientOpt{
		Addr: "localhost:6379",
	}
	client := asynq.NewClient(r)

	t1 := tasks.NewWelcomeEmailTask(42)
	t2 := tasks.NewWelcomeEmailTask(42)

	res, err := client.Enqueue(t1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result:%+v\n", res)

	res, err = client.Enqueue(t2, asynq.ProcessIn(24*time.Hour))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result:%+v\n", res)
}
