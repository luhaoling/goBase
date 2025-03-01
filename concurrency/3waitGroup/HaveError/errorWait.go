package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// don't execute
	go doSomething(100, &wg)
	go doSomething(120, &wg)
	go doSomething(130, &wg)
	go doSomething(140, &wg)
	go doSomething(140, &wg)
	go doSomething(140, &wg)
	go doSomething(140, &wg)
	go doSomething(140, &wg)
	go doSomething(140, &wg)
	go doSomething(140, &wg)

	wg.Wait()
	fmt.Println("over")
}

func doSomething(millisecs time.Duration, wg *sync.WaitGroup) {
	wg.Add(1)
	duration := millisecs * time.Millisecond
	time.Sleep(duration)

	fmt.Println("后台执行")
	wg.Done()

}

//
//func main() {
//	var wg sync.WaitGroup
//	wg.Add(10)
//	go doSomething(100, &wg)
//	go doSomething(120, &wg)
//	go doSomething(130, &wg)
//	go doSomething(140, &wg)
//	go doSomething(140, &wg)
//	go doSomething(140, &wg)
//	go doSomething(140, &wg)
//	go doSomething(140, &wg)
//	go doSomething(140, &wg)
//	go doSomething(140, &wg)
//
//	wg.Wait()
//	fmt.Println("over")
//}
//
//func doSomething(millisecs time.Duration, wg *sync.WaitGroup) {
//
//	duration := millisecs * time.Millisecond
//	time.Sleep(duration)
//
//	fmt.Println("后台执行")
//	wg.Done()
//
//}
