package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	GoSafe(func() {
		fmt.Println("nihao")
	})
	time.Sleep(1 * time.Second)
}

func GoSafe(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("goroutine panic recovered:%v", r)
			}
		}()
		fn()
	}()
}
