package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// go 实现乐观锁
// go 通常基于版本号或 CAS 机制实现乐观锁。乐观锁一般适用于读锁写少、冲突概率极低的场景
// 核心原理
// 无锁读取：读取数据时不加锁，同时记录数据的当前版本号或状态
// 条件更新：更新时检查版本号是否与读取时一致，一致则更新成功，然后版本号递增；不一致则失败（需重试或处理冲突）

type OptimisticData struct {
	Value   interface{}
	Version int64
}

// 获取当前值以及版本号
func (d *OptimisticData) Get() (interface{}, int64) {
	version := atomic.LoadInt64(&d.Version)
	return d.Value, version
}

// 更新值，传入版本号和新的值
func (d *OptimisticData) Update(expectedVersion int64, newValue interface{}) bool {
	// 对比版本号，如果版本号相同，则修改值，否则返回 false
	if atomic.CompareAndSwapInt64(&d.Version, expectedVersion, expectedVersion+1) {
		d.Value = newValue
		return true
	}
	return false
}

type Account struct {
	balance OptimisticData
}

func (a *Account) WithDraw(amount float64) error {
	for {
		currentBalance, currentVersion := a.balance.Get()
		current := currentBalance.(float64)
		if current < amount {
			return fmt.Errorf("余额不足")
		}
		newBalance := current - amount
		if a.balance.Update(currentVersion, newBalance) {
			return nil
		}

		// 版本冲突，根据业务进行下一步操作
	}
}

func main() {
	account := &Account{
		balance: OptimisticData{Value: 100.0},
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := account.WithDraw(10); err != nil {
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()
	fmt.Println("最终余额:", account.balance.Value)
}
