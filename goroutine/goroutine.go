package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	Usage()
}

// 启动的函数没有参数
func GoSafe(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("goroutine panic recovered:%v", r)
			}
		}()
		fn()
	}()
}

// 支持自定义错误处理
func GoSafeWithHandler(fn func(), panicHandler func(interface{})) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				if panicHandler != nil {
					panicHandler(r)
				} else {
					log.Printf("default panic handling:%v", r)
				}
			}
		}()
		fn()
	}()
}

func GoSafeWithHandlerUsage() {
	GoSafeWithHandler(func() {
		fmt.Println("nihao")
	}, func(i interface{}) {
		fmt.Println(i)
	})
	time.Sleep(1 * time.Second)
}

type SafeGo struct {
	wg      sync.WaitGroup
	ErrChan chan error
}

func NewSafeGo() *SafeGo {
	return &SafeGo{
		ErrChan: make(chan error, 10),
	}
}

func (sg *SafeGo) Go(fn func() error) {
	sg.wg.Add(1)
	go func() {
		defer sg.wg.Done()
		defer func() {
			if r := recover(); r != nil {
				sg.ErrChan <- fmt.Errorf("panic recovered:%v", r)
			}
		}()
		if err := fn(); err != nil {
			sg.ErrChan <- err
		}
	}()
}

func (sg *SafeGo) Wait() {
	go func() {
		sg.wg.Wait()
		close(sg.ErrChan)
	}()
}

func Usage() {
	sg := NewSafeGo()
	sg.Go(func() error {
		time.Sleep(5 * time.Second)
		return errors.New("常规错误示例")
	})

	sg.Go(func() error {
		time.Sleep(1 * time.Second)
		panic("紧急崩溃示例")
	})

	sg.Wait()

	for err := range sg.ErrChan {
		fmt.Printf("捕获错误:%v\n", err)
	}

}
