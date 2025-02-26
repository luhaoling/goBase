package main

import "fmt"

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

func main() {
	coms := buy(10)
	phones := build(coms)
	packs := pack(phones)
	for p := range packs {
		fmt.Println(p)
	}
}
