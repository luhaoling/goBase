package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type MyHandler struct {
	Title string
}

func (m *MyHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Printf("%s recv from %v, msg:%v\n", m.Title, msg.NSQDAddress, string(msg.Body))
	return nil
}

func initConsumer(topic string, channel string, address string) (err error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	c, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Printf("create consumer failed,err:%v\n", err)
		return
	}

	consumer := &MyHandler{
		Title: "沙河1号",
	}

	c.AddHandler(consumer)

	if err := c.ConnectToNSQD(address); err != nil {
		return err
	}

	return nil
}

func main() {
	err := initConsumer("topic_demo", "first", "127.0.0.1:4150")
	if err != nil {
		fmt.Printf("init comsumer failed,err:%v\n", err)
		return
	}
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	<-c
}
