package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.RWMutex

	go func() {
		time.Sleep(200 * time.Millisecond)
		mu.Lock()
		fmt.Println("Lock")
		time.Sleep(100 * time.Millisecond)
		mu.Unlock()
		fmt.Println("Unlock")
	}()
	go func() {
		factorial(&mu, 10)
	}()
	select {}
}

func factorial(m *sync.RWMutex, n int) int {
	if n < 1 {
		return 0
	}
	fmt.Println("RLock")
	m.RLock()
	defer func() {
		fmt.Println("RUnLock")
		m.RUnlock()
	}()
	time.Sleep(100 * time.Millisecond)
	return n * factorial(m, n-1)

}




