package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 要求：使用 goroutine 和 channel 实现一个函数，启动 3 个并发的任务（每个任务休眠随机时间后返回结果），
// 主协程等待所有子任务任务完成并收集结果。

//
//var ch1 chan int
//var ch2 chan int
//var ch3 chan int
//
//func main() {
//	ch1 = make(chan int)
//	ch2 = make(chan int)
//	ch3 = make(chan int)
//	fmt.Println(RunTasks())
//}
//
//func RunTasks() (result []int) {
//	go func() {
//		r := rand.Intn(10)
//		fmt.Println(r)
//		time.Sleep(time.Duration(r) * time.Second)
//		ch1 <- r
//	}()
//	go func() {
//		r := rand.Intn(10)
//		fmt.Println(r)
//		time.Sleep(time.Duration(r) * time.Second)
//		ch2 <- r
//	}()
//	go func() {
//		r := rand.Intn(10)
//		fmt.Println(r)
//		time.Sleep(time.Duration(r) * time.Second)
//		ch3 <- r
//	}()
//	result = append(result, <-ch1)
//	result = append(result, <-ch2)
//	result = append(result, <-ch3)
//	return
//}

// 优化
func RunTasks() []int {
	resultChan := make(chan int, 3)
	for i := 0; i < 3; i++ {
		go func(taskID int) {
			sleepTime := rand.Intn(10)
			time.Sleep(time.Duration(sleepTime) * time.Second)
			resultChan <- taskID
		}(i)
	}
	results := make([]int, 0, 3)
	for i := 0; i < 3; i++ {
		results = append(results, <-resultChan)
	}
	return results
}

func main() {
	fmt.Println(RunTasks())
}
