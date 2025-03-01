package main

import (
	"fmt"
	"sync"
)

//copy WaitGroup is not allow

type TestWaitGroup struct {
	Wait sync.WaitGroup
}

type test struct {
	a string
}

func main() {
	w := sync.WaitGroup{}
	w.Add(1)
	t := &TestWaitGroup{
		Wait: w,
	}
	t.Wait.Done()
	fmt.Println("finished")
}
