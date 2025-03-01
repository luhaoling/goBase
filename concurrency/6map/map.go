package main

import "fmt"

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
}
