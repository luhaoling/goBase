package limitRate

import (
	"sync"
	"time"
)

//https://blog.csdn.net/baidu_32452525/article/details/133647411

// 在一定时间间隔里，记录请求次数，当请求次数超过该时间限制时，就把计数器清零，然后重新计算

type CounterLimit struct {
	rate       int           // 每秒最大请求数量
	interval   time.Duration // 时间窗口大小
	counter    int           // 当前时间窗口内的请求数量
	lastUpdate time.Time     // 上次更新时间
	mu         sync.Mutex    // 互斥锁
}

func NewCounterLimiter(rate int, interval time.Duration) *CounterLimit {
	return &CounterLimit{
		rate:       rate,
		interval:   interval,
		counter:    0,
		lastUpdate: time.Now(),
	}
}

func (cl *CounterLimit) Allow() bool {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(cl.lastUpdate)

	// 如果超过时间窗口，重置计数器
	if elapsed >= cl.interval {
		cl.counter = 0
		cl.lastUpdate = now
	}

	// 判断当前时间窗口内的请i求数量是否超过限制
	if cl.counter < cl.rate {
		cl.counter++
		return true
	}
	return false
}
