package main

import (
	"fmt"
	"sync"
)

func main() {
	slice2()
}

func slice2() {
	sl := make([]int, 100)
	wg := sync.WaitGroup{}
	for index := 0; index < 100; index++ {
		k := index
		wg.Add(1)
		go func(num int) {
			sl[num] = num
			wg.Done()
		}(k)
	}
	wg.Wait()
	fmt.Printf("final len(sl)=%d cap(sl)=%d\n", len(sl), cap(sl))
	fmt.Println(sl)
}

func slice1() {
	sl := make([]int, 0)
	wg := sync.WaitGroup{}
	for index := 0; index < 100; index++ {
		k := index
		wg.Add(1)
		go func(num int) {
			sl = append(sl, num)
			wg.Done()
		}(k)
	}
	wg.Wait()
	fmt.Printf("final len(sl)=%d cap(sl)=%d\n", len(sl), cap(sl))
	fmt.Println(sl)
}
