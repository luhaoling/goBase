package main

import (
	"net/http"
	_ "net/http/pprof"
)

//to http://127.0.0.1:8888/debug/pprof/

//  直接访问 url 并添加参数来控制采集数据
//  curl "http://127.0.0.1:8888/debug/pprof/profile?seconds=30" > profile.pprof
// 在 Terminal 输入 go tool pprof -http :8889 profile.pprof
// go tool pprof -http :8889 http://127.0.0.1:8888/debug/pprof/profile

func main() {
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
