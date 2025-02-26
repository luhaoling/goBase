package main

import (
	"fmt"
	"testing"
	"time"
)

type MockIdGenerator struct{}

func (m *MockIdGenerator) Next() int64 {
	return time.Now().UnixNano()
}

type MockProducer struct{}

func (m *MockProducer) Send(msg Message) {
	fmt.Printf("message sent:%+v\n", msg)
}

func Test(t *testing.T) {
	idGen := &MockIdGenerator{}
	producer := &MockProducer{}
	handler := &RequestHandler{
		IdGenerator: idGen,
		Producer:    producer,
		Result:      make(map[int64]*Result),
		MyID:        1,
		Timeout:     2 * time.Second,
		MQ:          make(chan Message, 1),
	}

	request := Request{
		UserID:    123,
		LotteryID: 456,
	}

	response := handler.OnRequest(request)
	fmt.Printf("Response:%+v\n", response)

	handler.OnResult(Result{
		UUID: 12345,
		Flag: true,
	})

	response = handler.OnRequest(request)
	fmt.Printf("Second Response:%+v\n", response)
}
