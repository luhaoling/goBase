package main

import (
	"os"
	"runtime/pprof"
	"time"
)

// 执行完代码后
// 在 Terminal 输入> go tool pprof -http :8889 profile.pprof

func main() {
	// cpu采样
	f, err := os.Create("profile.pprof") // 创建在根目录
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()
	// 运行业务逻辑
	expensiveCPU()

	time.Sleep(30 * time.Second)
}

// 模拟高 CPU 消耗
func expensiveCPU() {
	for i := 0; i < 8; i++ {
		go func() {
			for {
				fib(35)
			}
		}()
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
