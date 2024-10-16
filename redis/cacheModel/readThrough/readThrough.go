package readThrough

import (
	"context"
	"fmt"
	"goBase/redis/common"
	"log"
	"strconv"
)

// 读穿透
// 写操作：先写数据库，后写缓存
// 读操作：
// 缓存中存在，直接返回缓存中的值
// 缓存中不存在（3 种处理方案）
// 1.从数据库中读取，并写入缓存，然后返回
// 2.从数据库中读取，读取到对应值后返回，异步写入缓存
// 3.直接返回错误响应或默认值，异步从数据库读取，然后写入缓存

// 使用场景
// 业务方对响应要求苛刻，可以考虑使用方案3。代价是业务方会收到错误响应或者默认值
// 如果缓存很慢或者是缓存大对象时以及把大对象序列化后存储在缓存中，可以考虑使用方案2

// 方案一实现

const (
	KEY = "readThrough:"
)

func NewReadThroughCache() *ReadThroughCache {
	return &ReadThroughCache{c: common.NewCache(), fn: func(ctx context.Context, key string) (any, error) {
		m := common.NewDatabase()
		id, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("key %s is not a number", key)
		}
		return m.Get(ctx, id)
	}}
}

type ReadThroughCache struct {
	c  common.Cache
	fn func(ctx context.Context, key string) (any, error) // 读数据库操作
}

func (r *ReadThroughCache) Get(ctx context.Context, key string) (any, error) {
	val, err := r.c.Get(ctx, KEY+key)
	if err != nil {
		log.Println(err)
	}
	if val == nil {
		val, err = r.fn(ctx, key)
		if err != nil {
			return nil, err
		}
		err = r.c.Set(context.Background(), KEY+key, val)
		if err != nil {
			log.Println("set cache err:", err)
		}
	}
	return val, nil
}
