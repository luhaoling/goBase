package main

import (
	"net/http"
	_ "net/http/pprof"
)

// // 访问对应端口并触发数据采集
// // curl "http://localhost:8888/debug/pprof/trace?seconds=30" > trace.out
// 方法1：
func main() {
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
