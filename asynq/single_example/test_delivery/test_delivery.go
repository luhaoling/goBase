package test_delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
)

const (
	TypeEmailDelivery = "email:deliver"
)

type EmailDeliveryPayload struct {
	UserID     int
	TemplateID string
	DetaStr    string
}

// 包装消息
func NewEmailDeliveryTask(userID int, tmp1ID, dataStr string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailDeliveryPayload{
		UserID:     userID,
		TemplateID: tmp1ID,
		DetaStr:    dataStr,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

// 拆解消息并处理
func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed:%v:%w", err, asynq.SkipRetry)
	}
	log.Printf("Sending Email to User:user_id=%d,template_id=%s data_str:%s", p.UserID, p.TemplateID, p.DetaStr)
	return nil
}
