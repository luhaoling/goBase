package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	two()
}

// 协程取消与退出
func one() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("子协程退出")
				return
			default:
				fmt.Println("nihao")
			}
		}
	}(ctx)
	time.Sleep(5 * time.Second)
	cancel()
}

// 网络请求超时
func two() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	time.Sleep(4 * time.Second)
	defer cancel()
	go func(ctx context.Context) {
		fmt.Println("查询中")
	}(ctx)
}

// 认证信息传递
func authMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := "token"
		ctx := context.WithValue(r.Context(), "token", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
