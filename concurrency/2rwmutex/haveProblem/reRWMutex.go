package main

import (
	"fmt"
	"sync"
)

func main() {
	l := &sync.RWMutex{}
	l.Lock()
	foo(l)
}

func foo(l *sync.RWMutex) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}

func bar(l *sync.RWMutex) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}
