package main

import (
	"fmt"
	"sync"
	"time"
)

type IdGenerator interface {
	Next() int64
}

type producer interface {
	Send(msg Message)
}

type Message struct {
	Request Request
	UUID    int64
	MyID    int64
}

type Result struct {
	UUID int64
	Flag bool
}

func (r *Result) Success() bool {
	return r.Flag
}

type Request struct {
	UserID    int64
	LotteryID int64
}

type Response struct {
	Success bool
	Message string
}

func SuccessResponse() Response {
	return Response{Success: true, Message: "秒杀成功"}
}

func TimeoutResponse() Response {
	return Response{Success: false, Message: "超时"}
}

func FailResponse() Response {
	return Response{Success: false, Message: "秒杀失败"}
}

type RequestHandler struct {
	IdGenerator IdGenerator
	Producer    producer
	Result      map[int64]*Result
	Mutexes     sync.Map
	MyID        int64
	Timeout     time.Duration
	MQ          chan Message
}

func (h *RequestHandler) OnRequest(request Request) Response {
	uuid := h.IdGenerator.Next()

	// 创建消息
	msg := h.ComposeMsg(request, uuid)

	mutex := &sync.Mutex{}

	h.Mutexes.Store(uuid, mutex)

	h.Send(msg)

	done := make(chan bool, 1)
	go func() {
		msg := <-h.MQ
		fmt.Printf("后端处理:%v", msg)
		time.Sleep(h.Timeout)
		h.Result[uuid] = &Result{
			UUID: msg.UUID,
			Flag: true,
		}
		done <- true
	}()

	select {
	case <-done:
		result, _ := h.Result[uuid]
		if result != nil && result.Success() {
			return SuccessResponse()
		}
	case <-time.After(2 * h.Timeout):
		return TimeoutResponse()
	}

	h.Mutexes.Delete(uuid)
	return FailResponse()
}

func (h *RequestHandler) Send(message Message) {
	h.MQ <- message
}

func (h *RequestHandler) OnResult(result Result) {
	value, ok := h.Mutexes.Load(result.UUID)
	if ok {
		mutex := value.(*sync.Mutex)
		h.Result[result.UUID] = &result
		mutex.Lock()
		mutex.Unlock()
	}
}

func (h *RequestHandler) ComposeMsg(request Request, uuid int64) Message {
	return Message{
		Request: request,
		UUID:    uuid,
		MyID:    h.MyID,
	}
}
