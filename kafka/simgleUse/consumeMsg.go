package main

import (
	"context"
	"fmt"
	"goBase/kafka/simgleUse/consumer"
	"os"
	"os/signal"
	"syscall"
)

func listenSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	fmt.Printf("收到信号:%s", sig.String())
	if consumer.Consumer != nil {
		consumer.Consumer.Close()
	}
	os.Exit(0)
}

func main() {
	ctx := context.Background()
	go consumer.ReadMessage(ctx)
	listenSignal()
}
