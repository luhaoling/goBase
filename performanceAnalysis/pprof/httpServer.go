package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

// ab 工具压测（模拟线上用户，对 http 服务发起请求压测）
// > ab -n 30000 -c 2 http://127.0.0.1:8888/

// 如果在 ab 测试中发现内存占用较高，可以使用 pprof 采集单机内存报告并生成内存火焰图
// curl "http://127.0.0.1:8888/debug/pprof/heap?seconds=30" > heap.pprof

// go tool pprof -http :8889 heap.pprof

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
	n := 1000
	slices := make([]int, 0)
	for i := 0; i < n; i++ {
		slices = append(slices, i)
	}
	return slices
}

func getRawSet(slices []int) map[int]bool {
	set := make(map[int]bool, 0)
	for _, item := range slices {
		set[item] = true
	}
	return set
}
