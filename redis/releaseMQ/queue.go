package releaseMQ

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
)

// from: https://learnku.com/articles/87043

const (
	HashSuffix = ":hash" // redis 键后缀，用于哈希表
	SetSuffix  = ":set"  // redis 键后缀，用于集合
)

var once sync.Once

type Queue struct {
	ctx context.Context

	redis *redis.Client
	topic string

	producer *producer
	consumer *consumer
}

func NewQueue(ctx context.Context, redis *redis.Client, opts ...Option) *Queue {
	var queue *Queue

	once.Do(func() {
		defaultOptions := Options{
			topic:   "topic",
			handler: defaultHander,
		}

		for _, apply := range opts {
			apply(&defaultOptions)
		}

		queue = &Queue{
			ctx:      ctx,
			redis:    redis,
			topic:    defaultOptions.topic,
			producer: NewProducer(ctx),
			consumer: NewConsumer(ctx, defaultOptions.Handler),
		}

	})

	return queue
}

func (q *Queue) Start() {
	go q.consumer.listen(q.redis, q.topic)
}

func (q *Queue) Publish(msg *Message) (int64, error) {
	return q.producer.publish(q.redis, q.topic, msg)
}

