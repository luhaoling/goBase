package main

import (
	"fmt"
	"sync"
	"time"
)

// pipelineModel
// 由一道道工序组成一条流水线
func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- fmt.Sprintf("配件:%d", i)
		}
	}()
	return out
}

func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "组装(" + c + ")"
		}
	}()
	return out
}

func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "打包(" + c + ")"
		}
	}()
	return out
}

//func main() {
//	coms := buy(10)
//	phones := build(coms)
//	packs := pack(phones)
//	for p := range packs {
//		fmt.Println(p)
//	}
//}

// 扇入、扇出模式
// 在某一道工序中增加人手，提高工作效率
func merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup

	out := make(chan string)
	p := func(in <-chan string) {
		defer wg.Done()
		for c := range in {
			out <- c
		}
	}
	wg.Add(len(ins))
	for _, cs := range ins {
		go p(cs)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

//func main() {
//	coms := buy(100)
//	phones1 := build(coms)
//	phones2 := build(coms)
//	phones3 := build(coms)
//	phones := merge(phones1, phones2, phones3)
//	packs := pack(phones)
//	for p := range packs {
//		fmt.Println(p)
//	}
//}

// futures 模式
// Futures 模式下的协程和普通协程最大的区别是可以返回结果，而这个结果会在未来的某个时间点使用。
// 所以在未来获取这个结果的操作必须是一个阻塞的操作，要一直等到获取结果为止。
func washVegetables() <-chan string {
	vegatables := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		vegatables <- "洗好的菜"
	}()
	return vegatables
}

// 烧水
func boilWater() <-chan string {
	water := make(chan string)
	go func() {
		time.Sleep(8 * time.Second)
		water <- "烧开的水"
	}()
	return water
}

func main() {
	vegetalbesCh := washVegetables()
	waterCh := boilWater()
	fmt.Println("已经安排洗菜和烧水了，我先眯一会")
	time.Sleep(2 * time.Second)
	fmt.Println("要做火锅了，菜和水好了吗？")
	vegetalbes := <-vegetalbesCh
	waster := <-waterCh
	fmt.Println("准备好了，可以做火锅了.", vegetalbes, waster)
}
