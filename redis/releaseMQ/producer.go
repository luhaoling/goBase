package queue

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// Producer：消息生产者，用于向队列发布消息

type producer struct {
	ctx context.Context
}

func NewProducer(ctx context.Context) *producer {
	return &producer{
		ctx: ctx,
	}
}

// Publish 将消息发布到 Redis 中
func (p *producer) publish(redisClient *redis.Client, topic string, msg *Message) (int64, error) {
	z := &redis.Z{
		Score:  msg.GetScore(), // 即消费时间
		Member: msg.GetId(),    // 消息唯一标识
	}

	// 将消息写入有序集合
	setKey := topic + SetSuffix
	n, err := redisClient.ZAdd(p.ctx, setKey, z).Result()
	if err != nil {
		return n, err
	}

	// 将消息写入哈希表
	hashKey := topic + HashSuffix
	return redisClient.HSet(p.ctx, hashKey, msg.GetId(), msg).Result()
}
