package main

import (
	"goBase/asynq/single_example/test_delivery/client"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		client.EmailDeliveryTaskAdd(i)
		time.Sleep(time.Second * 3)
	}
}
