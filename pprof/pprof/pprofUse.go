package main

import (
	"net/http"
	_ "net/http/pprof"
)

//to http://127.0.0.1:8888/debug/pprof/

func main() {
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
