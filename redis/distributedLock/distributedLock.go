package distributedLock

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

// 如何实现一个分布式锁
// 主要用到的方法 redis的 setnx 和 del
// 主要内容有:
// 1. 获取锁
// 2. 释放锁(要保证只能释放自己的锁)
// 3. 续期
// 4. 获取锁的超时时间
// 5. 获取锁的等待时间
// 6. 超时重试

// 主要方法：获取锁和释放锁
const (
	LOCK_EXPIRE = 60 * time.Second
	DISLOCK     = "dis_lock:lock"
)

type DistributedLock interface {
	TryLock(ctx context.Context) error
	UnLock(ctx context.Context) error
}

type disLock struct {
	Uid   string
	redis *redis.Client
}

func (d *disLock) TryLock(ctx context.Context) error {
	var flag bool
	var err error
	for !flag {
		flag, err = d.redis.SetNX(ctx, DISLOCK, d.Uid, LOCK_EXPIRE).Result()
		if err != nil {
			fmt.Errorf("get lock error:%v", err)
		}
	}
	log.Printf("get lock success")
	return nil
}

func (d *disLock) UnLock(ctx context.Context) error {
	_, err := d.redis.Get(ctx, DISLOCK).Result()
	if err != nil {
		return fmt.Errorf("get lock error:%v", err)
	}
	return nil
}

func NewDistributedLock(uid string, redis *redis.Client) DistributedLock {
	return &disLock{
		Uid:   uid,
		redis: redis,
	}
}
