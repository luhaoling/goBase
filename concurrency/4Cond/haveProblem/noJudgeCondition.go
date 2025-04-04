package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	var ready int

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
			c.L.Lock()
			ready++
			c.L.Unlock()
			log.Printf("运动员%d 已准备就绪\n", i)
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	// 不判断条件，导致所有运动员还没准备好就开始比赛了
	//for ready != 10 {
	c.Wait()
	//}
	c.L.Unlock()
	log.Println("所有运动员准备就绪，比赛开始，3，2，1，GO！")

}
