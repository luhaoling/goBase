package main

import "sync"

type Once struct {
	m sync.Mutex
}

func (o *Once) doSlow() {
	o.m.Lock()
	defer o.m.Unlock()

	*o = Once{}
}

func main() {
	var once Once
	once.doSlow()
}
