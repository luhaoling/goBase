package queue

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type producer struct {
	ctx context.Context
}

func NewProducer(ctx context.Context) *producer {
	return &producer{
		ctx: ctx,
	}
}

func (p *producer) publish(redisClient *redis.Client, topic string, msg *Message) (int64, error) {
	z := &redis.Z{
		Score:  msg.GetScore(),
		Member: msg.GetId(),
	}

	setKey := topic + SetSuffix
	n, err := redisClient.ZAdd(p.ctx, setKey, z).Result()
	if err != nil {
		return n, err
	}

	hashKey := topic + HashSuffix
	return redisClient.HSet(p.ctx, hashKey, msg.GetId(), msg).Result()
}
