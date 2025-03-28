package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	c, cancel := context.WithTimeout(ctx, 10*time.Second)
	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	fmt.Println(struct{}{} == <-c.Done())

	ch := make(chan struct{})
	go func() {
		time.Sleep(3 * time.Second)
		close(ch)
	}()
	fmt.Println(struct{}{} == <-ch)

	c1 := context.WithValue(ctx, "key", "value")
	fmt.Println(c1.Value("key"))
}
