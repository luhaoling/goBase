package common

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) (any, error)
	Set(ctx context.Context, key string, val any) error
}

type cache struct {
	redis *redis.Client
}

func NewCache() Cache {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       6,  // use default DB
	})
	_, err := redis.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return &cache{
		redis: redis,
	}
}

func (c *cache) Get(ctx context.Context, key string) (any, error) {
	val, err := c.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("%s in redis is nil", key)
	}
	var user User
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		return nil, fmt.Errorf("json %s unmarshal err:%v", val, err)
	}
	return val, nil

}

func (c *cache) Set(ctx context.Context, key string, val any) error {
	data, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("json marshal err:%v", err)
	}
	return c.redis.Set(ctx, key, data, 60*time.Second).Err()
}
