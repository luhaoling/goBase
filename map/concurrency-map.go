package main

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
)

func main() {
	m := cmap.New[int]()

	m.Set("counter", 0)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				m.Upsert("counter", 1, func(exist bool, oldVal int, newVal int) int {
					return oldVal + 1
				})
			}
		}()
	}
	fmt.Println(m.Get("counter"))
}
