package queue

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"time"
)

// Message：消息实体，定义了消息的结构和序列化方法

type Message struct {
	Id          string      `json:"id"`
	CreateTime  time.Time   `json:"createTime"`  // 消息生产时间
	ConsumeTime time.Time   `json:"consumeTime"` // 要被消费的时间
	Body        interface{} `json:"body"`        // 消息体
}

func NewMessage(id string, consumeTime time.Time, body interface{}) *Message {
	if id == "" {
		id = uuid.NewV4().String()
	}
	return &Message{
		Id:          id,
		CreateTime:  time.Now(),
		ConsumeTime: consumeTime,
		Body:        body,
	}
}

// GetScore 返回消息的分数
func (m *Message) GetScore() float64 {
	return float64(m.ConsumeTime.Unix())
}

func (m *Message) GetId() string {
	return m.Id
}

// MarshalBinary 用于消息的序列化(将消息结构体序列化为二进制数据)
func (m *Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

// UnmarshalBinary 用于消息的反序列化(将二进制数据反序列化为消息结构体)
func (m *Message) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}
