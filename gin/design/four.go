package main

import (
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

// 要求：使用 Gin 框架实现一个全局计数器，满足以下条件：
// 定义 GET 方法返回当前计数器值
// 定义 Post 方法将计数器 +1，返回更新后的值
// 添加限流中间件，限制每秒最多 5 次 请求（超出返回 429 状态码）
// 保证计数器的并发安全

var (
	counter     int
	mu          sync.Mutex
	rateLimiter = NewTokenBucket(5, 5)
)

func main() {
	rateLimiter.Start()
	defer rateLimiter.Stop()
	router := gin.Default()

	router.Use(middle())
	router.GET("/counter", getValue)
	router.POST("/counter", AddValue)
	router.Run(":8080")
}

func getValue(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	c.JSON(200, gin.H{
		"counter": counter,
	})
}

func AddValue(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	counter++
	c.JSON(200, gin.H{
		"counter": counter,
	})
}

func middle() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !rateLimiter.Allow() {
			c.JSON(429, "rate limit")
			c.Abort()
			return
		}
		c.Next()
	}
}

// release a tokenBucket
type tokenBucket struct {
	limitRate int
	tokenChan chan struct{}
	cap       int
	muLock    *sync.Mutex
	stop      bool
}

func NewTokenBucket(limitRate int, cap int) *tokenBucket {
	if cap < 1 {
		panic("token bucket cap must be large 1")
	}

	b := &tokenBucket{
		tokenChan: make(chan struct{}, cap),
		limitRate: limitRate,
		cap:       cap,
		muLock:    new(sync.Mutex),
		stop:      false,
	}
	for i := 0; i < cap; i++ {
		b.tokenChan <- struct{}{}
	}
	return b
}

func (b *tokenBucket) Start() {
	go b.produce()
}

func (b *tokenBucket) produce() {
	ticker := time.NewTicker(time.Second / time.Duration(b.limitRate))
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			b.muLock.Lock()
			if !b.stop && len(b.tokenChan) < b.cap {
				b.tokenChan <- struct{}{}
			}
			b.muLock.Lock()
		default:

		}
	}
}

func (b *tokenBucket) Allow() bool {
	select {
	case <-b.tokenChan:
		return true
	default:
		return false
	}

}

func (b *tokenBucket) Stop() {
	b.muLock.Lock()
	defer b.muLock.Unlock()
	b.stop = true
}
