package main

import (
	"fmt"
	_ "github.com/mkevac/debugcharts" // 添加后可以查看实时图表数据
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
)

// this code from https://blog.csdn.net/weixin_45565886/article/details/132090495

// to http://localhost/debug/pprof/
// to http://localhost:80/debug/charts/
// 查看内存使用情况 （控制台输入）：go tool pprof  http://localhost:80/debug/pprof/heap
// 控制台交互模式下输入
//> help可以查看使用说明
//
//> top 可以查看前10个的内存分配情况
//
//> tree 以树状显示
//
//> png 以图片格式输出
//
//> svg 生成浏览器可以识别的svg文件

func createPprof() {
	f, err := os.Create("profile.pprof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()
	//expensiveCPU()
}

// Terminal 输入> go tool pprof -http :8889 profile.pprof
func main() {
	f, err := os.Create("heap.pprof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	runtime.GC()
	expensiveMem()
	if err := pprof.WriteHeapProfile(f); err != nil {
		panic(err)
	}
}

func expensiveMem() {
	ch := make(map[int]int)
	for i := 0; i < 10000; i++ {
		ch[i] = i
	}
	m := make([]int, 10000)
	anotherExpensiveMem(m)
}

func anotherExpensiveMem(arr []int) {
	for i, _ := range arr {
		fmt.Println(i)
	}
}
