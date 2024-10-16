package main

import (
	"time"
)

var m map[int]int

func main() {
	m = make(map[int]int)
	go m1()
	go m2()
	time.Sleep(5 * time.Second)
	for k, v := range m {
		println(k, v)
	}
}

func m1() {

	for i := 0; i < 100; i++ {
		m[i] = i
	}

}

func m2() {
	for i := 0; i < 100; i++ {
		m[i+100] = 100 + i
	}
}
