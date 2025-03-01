package main

import "sync"

func main() {
	var counter CounterM
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {

		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				counter.Incr()
			}
		}()
	}
	wg.Wait()
	println(counter.Count())
}

type CounterM struct {
	CounterType int
	Name        string

	mu    sync.Mutex
	count uint64
}

func (c *CounterM) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *CounterM) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
