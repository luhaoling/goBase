package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	count uint64
}

func main() {
	var wg sync.WaitGroup
	var counter Counter
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go worker(&counter, &wg)
	}
	wg.Wait()
	fmt.Println(counter.Count())
}
func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func worker(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	c.Incr()
}

func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}
