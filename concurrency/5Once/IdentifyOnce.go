package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

// 既可以返回当前调用 Do 方法是否正确完成，还可以在初始化失败后调用 Do 方法再次尝试初始化，直到成功

func (o *Once) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}
	return o.slowDo(f)
}

func (o *Once) slowDo(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 {
		err = f()
		if err == nil {
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}

// 判断是否初始化过
func (o *Once) Done() bool {
	return atomic.LoadUint32(&o.done) == 1
}

func main() {
	var flag Once
	fmt.Println(flag.Done())

	err := flag.Do(func() error {
		time.Sleep(time.Second)
		return nil
	})
	if err != nil {
		fmt.Println("failed")
	}

	fmt.Println(flag.Done())
}
