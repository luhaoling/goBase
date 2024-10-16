package main

import (
	"github.com/hibiken/asynq"
	"goBase/asynq/synchronized/tasks"
	"log"
)

type User struct {
	UserID int
}

func main() {
	r := asynq.RedisClientOpt{
		Addr: "localhost:6379",
	}
	srv := asynq.NewServer(r, asynq.Config{
		Concurrency: 10,
	})

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.WelcomeEmail, tasks.HandleWelcomeEmailTaskFaild)
	mux.HandleFunc(tasks.ReminderEmail, tasks.HandleReminderEmailTask)
	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
