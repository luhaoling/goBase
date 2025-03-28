package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
	"time"
)

// Consumer：消息消费者：用于从队列接受并处理消息

// handlerFunc 用于处理消息的函数类型
type handlerFunc func(msg Message)

// 默认处理函数，打印消息
func defaultHandler(msg Message) {
	fmt.Println(msg)
}

// consumer 包含消费者相关的字段和方法
type consumer struct {
	ctx      context.Context
	duration time.Duration // 消费速率
	ch       chan []string
	handler  handlerFunc // 消息处理函数
}

// NewConsumer 创建唯一消费者实例
func NewConsumer(ctx context.Context, handler handlerFunc) *consumer {
	return &consumer{
		ctx:      ctx,
		duration: time.Second,
		ch:       make(chan []string, 1000),
		handler:  handler,
	}
}

// listen 方法用于监听消息队列中的消息并处理
func (c *consumer) listen(redisClient *redis.Client, topic string) {
	// 从哈希表中获取数据并处理
	go func() {
		for {
			select {
			case ret := <-c.ch:
				// 从哈希表中批量获取数据信息
				key := topic + HashSuffix
				result, err := redisClient.HMGet(c.ctx, key, ret...).Result()
				if err != nil {
					log.Println(err)
				}

				// 从哈希表中删除消息
				if len(result) > 0 {
					redisClient.HDel(c.ctx, key, ret...)
				}

				msg := Message{}
				for _, v := range result {
					// 由于哈希表和有序集合中的操作不是原子操作，可能会出现删除了集合中的数据但哈希表中的数据未删除的情况
					if v == nil {
						continue
					}
					str := v.(string)
					json.Unmarshal([]byte(str), &msg)
					go c.handler(msg)
				}
			}
		}
	}()

	// 定时器用于定时获取消息并处理
	ticker := time.NewTicker(c.duration)
	defer ticker.Stop()
	for {
		select {
		case <-c.ctx.Done(): // 上下文取消，退出监听
			log.Println("consumer quit:", c.ctx.Err())
			return
		case <-ticker.C: // 定时获取消息
			min := strconv.Itoa(0)
			max := strconv.Itoa(int(time.Now().Unix()))
			opt := &redis.ZRangeBy{
				Min: min,
				Max: max,
			}

			key := topic + SetSuffix
			result, err := redisClient.ZRangeByScore(c.ctx, key, opt).Result()
			if err != nil {
				log.Fatal(err)
				return
			}
			for _, v := range result {
				fmt.Println(v)
			}

			// 获取到数据
			if len(result) > 0 {
				// 从有序集合中移除数据
				redisClient.ZRemRangeByScore(c.ctx, key, min, max)
				// 写入通道，进行哈希表处理
				c.ch <- result
			}
		}
	}
}
