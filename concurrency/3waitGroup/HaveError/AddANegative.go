package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Add(-10)
	wg.Add(-1)
}
