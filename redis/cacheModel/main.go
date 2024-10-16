package main

import (
	"context"
	"fmt"
	"goBase/redis/cacheModel/readThrough"
)

func main() {
	r := readThrough.NewReadThroughCache()
	val, err := r.Get(context.Background(), "1")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("result:%+v", val)

}
