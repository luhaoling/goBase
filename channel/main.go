package main

import (
	"fmt"
	"time"
)

// 一共有三个协程，分别按顺序打印 1,2,3，要求打印100次

// 分配定义三个 channel，用于通知对应的 goroutine 打印对应的数字
// 在 for 的一次循环中开启三个协程，这三个协程都会阻塞，直到收到对应的运行信号
// 执行 ch1<-1 这个命令时，通知 print1 打印 1，print1 打印完 1 后，通知 print2 打印 2，
// print2 打印完 2 后，通知 print3 打印 3。由此按顺序打印 1，2，3。由 for 循环控制打印 100次

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	for i := 0; i < 100; i++ {
		go print1(ch1, ch2)
		go print2(ch2, ch3)
		go print3(ch3)
		ch1 <- 1
	}

	time.Sleep(10 * time.Second)
}

func print1(ch1, ch2 chan int) {
	<-ch1
	fmt.Println(1)
	ch2 <- 1
}

func print2(ch2, ch3 chan int) {
	<-ch2
	fmt.Println(2)
	ch3 <- 1
}

func print3(ch3 chan int) {
	<-ch3
	fmt.Println(3)
}
