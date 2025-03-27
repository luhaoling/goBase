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

type handlerFunc func(msg Message)

func defaultHandler(msg Message) {
	fmt.Println(msg)
}

type consumer struct {
	ctx      context.Context
	duration time.Duration
	ch       chan []string
	handler  handlerFunc
}

func NewConsumer(ctx context.Context, handler handlerFunc) *consumer {
	return &consumer{
		ctx:      ctx,
		duration: time.Second,
		ch:       make(chan []string, 1000),
		handler:  handler,
	}
}

func (c *consumer) listen(redisClient *redis.Client, topic string) {
	go func() {
		for {
			select {
			case ret := <-c.ch:
				key := topic + HashSuffix
				result, err := redisClient.HMGet(c.ctx, key, ret...).Result()
				if err != nil {
					log.Println(err)
				}

				if len(result) > 0 {
					redisClient.HDel(c.ctx, key, ret...)
				}

				msg := Message{}
				for _, v := range result {
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

	ticker := time.NewTicker(c.duration)
	defer ticker.Stop()
	for {
		select {
		case <-c.ctx.Done():
			log.Println("consumer quit:", c.ctx.Err())
			return
		case <-ticker.C:
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

			if len(result) > 0 {
				redisClient.ZRemRangeByScore(c.ctx, key, min, max)
				c.ch <- result
			}
		}
	}
}
