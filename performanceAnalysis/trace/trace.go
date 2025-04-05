package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/trace"
)

// 访问对应端口并触发数据采集
// curl "http://localhost:8888/debug/pprof/trace?seconds=30" > trace.out
// go tool trace trace.out
func main() {
	f, err := os.Create("trace1.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace:%v", err)
	}
	defer trace.Stop()

	err = http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
