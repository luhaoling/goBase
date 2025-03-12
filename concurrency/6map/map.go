package main

import (
	"fmt"
	"time"
)

type mapKey struct {
	key int
}

func main() {
	var m = make(map[mapKey]string)
	var key = mapKey{key: 1}
	m[key] = "hello"
	fmt.Printf("m[key]%s\n", m[key])

	key.key = 1000
	fmt.Printf("m[key]%s\n", m[key])

	map1 := make(map[int]int)
	for i := 0; i < 10; i++ {
		map1[i] = i
	}

	for i := 0; i < 10000; i++ {
		go func() {
			//map1[i]=i
			for i := 0; i < 10; i++ {

				fmt.Println(map1[i])
			}
		}()
	}
	time.Sleep(5 * time.Second)
}
