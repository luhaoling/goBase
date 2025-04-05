package main

import (
	"os"
	"runtime/pprof"
	"time"
)

// 执行程序后生成
// go tool pprof -http=:8889 heap.pprof

var keep [][]int // 保留堆对象引用，避免被 GC

func main() {
	go expensiveHeap()

	time.Sleep(5 * time.Second) // 等 goroutine 跑起来

	f, err := os.Create("heap.pprof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 不手动 GC，让活跃对象保留
	if err := pprof.WriteHeapProfile(f); err != nil {
		panic(err)
	}

	time.Sleep(30 * time.Second)
}

func expensiveHeap() {
	for i := 0; i < 8; i++ {
		go func() {
			for {
				// 制造堆内存分配
				data := make([]int, 10000)
				keep = append(keep, data) // 保留引用，防止 GC
				time.Sleep(10 * time.Millisecond)
			}
		}()
	}
}
