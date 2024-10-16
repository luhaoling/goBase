package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
)

const (
	WelcomeEmail  = "email:welcome"
	ReminderEmail = "email:reminder"
)

type User struct {
	UserID int `json:user_id`
}

func NewWelcomeEmailTask(id int) *asynq.Task {
	payload, err := json.Marshal(map[string]interface{}{"user_id": id})
	if err != nil {
		log.Fatal(err)
	}
	return asynq.NewTask(WelcomeEmail, payload)
}

func NewReminderEmailTask(id int) *asynq.Task {
	payload, err := json.Marshal(map[string]interface{}{"user_id": id})
	if err != nil {
		log.Fatal(err)
	}
	return asynq.NewTask(ReminderEmail, payload)
}

func HandleWelcomeEmailTask(ctx context.Context, t *asynq.Task) error {
	var m map[string]interface{}
	err := json.Unmarshal(t.Payload(), &m)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("sending welcome email to user %v", m["user_id"])
	return nil
}

func HandleReminderEmailTask(ctx context.Context, t *asynq.Task) error {
	var m map[string]interface{}
	err := json.Unmarshal(t.Payload(), &m)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("sending welcome email to user %v", m["user_id"])
	return nil
}

func HandleWelcomeEmailTaskFaild(ctx context.Context, t *asynq.Task) error {
	var m map[string]interface{}
	err := json.Unmarshal(t.Payload(), &m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("sending welcome email to user %v", m["user_id"])
	return fmt.Errorf("could not send email to the user")
}
