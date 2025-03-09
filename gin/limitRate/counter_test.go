package limitRate

import (
	"fmt"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	limiter := NewCounterLimiter(10, time.Second) // 每秒最多处理 10 个请求

	for i := 1; i <= 15; i++ {
		if limiter.Allow() {
			fmt.Println("Request", i, "allowed")
		} else {
			fmt.Println("Request", i, "rejected")
		}
		time.Sleep(20 * time.Millisecond)
	}
}
