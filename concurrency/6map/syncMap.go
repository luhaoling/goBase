package main

import (
	"fmt"
	"sync"
)

func main() {
	sym := sync.Map{}
	sym.Store("key", "map1")
	fmt.Println(sym.Load("key"))

}
