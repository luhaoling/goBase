package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

// 项目解析1：
// 第一步：运行项目
// go run httpServer.go

// 第二步：压测
// ab 工具压测（模拟线上用户，对 http 服务发起请求压测）
// > ab -n 30000 -c 2 http://127.0.0.1:8888/

// 第三步：采集 heap profile 文件
// 方法1：(windows 上执行会出错)
// 如果在 ab 测试中发现内存占用较高，可以使用 pprof 采集单机内存报告并生成内存火焰图
// curl "http://127.0.0.1:8888/debug/pprof/heap?seconds=30" > heap.pprof
// go tool pprof -http :8889 heap.pprof

// 方法2：
// 跳过文件生成步骤，直接通过 go tool pprof 连接服务端
// go tool pprof -http :8889 http://127.0.0.1:8888/debug/pprof/heap

// 项目解析2
// 步骤1：编译程序
// go build -o myserver.exe main.go
// 步骤2：打开一个新终端，运行服务：
// ./myserver.exe
// 步骤3：ab 压测
// ab -n 30000 -c 2 http://127.0.0.1:8888/
// 步骤3：在浏览器访问，如果能打开/下载，说明服务 ok
//	http://127.0.0.1:8888/debug/pprof/heap
// 步骤4：在另外一个终端执行
// go tool pprof -http=:8889 myserver.exe http://127.0.0.1:8888/debug/pprof/heap

func Handler(w http.ResponseWriter, r *http.Request) {
	slices := getRawSlices()
	getRawSet(slices)
	w.Header().Set("Context-Type", "text/plain")
	fmt.Fprintln(w, "hh")
}

func main() {
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

func getRawSlices() []int {
	n := 100000
	slices := make([]int, 0, n)
	for i := 0; i < n; i++ {
		slices = append(slices, i)
	}
	return slices
}

func getRawSet(slices []int) map[int]bool {
	set := make(map[int]bool, len(slices))
	for _, item := range slices {
		set[item] = true
	}
	return set
}
