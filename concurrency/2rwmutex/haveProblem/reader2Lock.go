package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.RWMutex

	go func() {
		mu.RLock()
		fmt.Println("Reader:Acquired RLock")
		defer mu.RUnlock()
		time.Sleep(1 * time.Second)

		fmt.Println("Reader:trying to acquire lock (write lock)")
		mu.Lock()
		fmt.Println("Reader:acquired lock")
		mu.Unlock()
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Writer:trying to acquire lock (write lock)")
		mu.Lock()
		fmt.Println("Writer:Acquire Lock")
		mu.Unlock()
	}()

	select {}
}