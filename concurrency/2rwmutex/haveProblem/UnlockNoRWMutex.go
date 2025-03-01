package main

import "sync"

func main() {
	var rw sync.RWMutex
	rw.RUnlock()

}
