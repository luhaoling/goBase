package main

import (
	"sync"
)

const SHARD_COUNT = 32

const PRIME32 uint32 = 16777619

// 分片锁的原理，定义多个带有 map 和 读写锁的结构体，组成一个数组。实例化这个数组后，每次往里面插入数据时，
// 先根据 key 计算它应该落在哪个位置上，然后获取这个位置的读写锁，并往里面插入数据，然后再解锁。查找操作也同理，
// 先找到对应的 key 的位置，然后加锁，获取值，然后再解锁。
type ConcurrentMap []*ConcurrentMapShared

type ConcurrentMapShared struct {
	items map[string]interface{}
	sync.RWMutex
}

func New() ConcurrentMap {
	m := make(ConcurrentMap, SHARD_COUNT)
	for i := 0; i < SHARD_COUNT; i++ {
		m[i] = &ConcurrentMapShared{
			items: make(map[string]interface{}),
		}
	}
	return m
}

func (m ConcurrentMap) GetShard(key string) *ConcurrentMapShared {
	return m[uint(fnv32(key))%uint(SHARD_COUNT)]
}

// FNV hash
func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	for i := 0; i < len(key); i++ {
		hash *= PRIME32
		hash ^= uint32(key[i])
	}
	return hash
}

func (m ConcurrentMap) Set(key string, value interface{}) {
	shard := m.GetShard(key)
	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

func (m ConcurrentMap) Get(key string) (interface{}, bool) {
	shard := m.GetShard(key)
	shard.RLock()
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

// 统计当前分段 map 中 item 的个数
func (m ConcurrentMap) Count() int {
	count := 0
	for i := 0; i < SHARD_COUNT; i++ {
		shard := m[i]
		shard.RLock()
		count += len(shard.items)
		shard.RUnlock()
	}
	return count
}

// 获取所有的 key
func (m ConcurrentMap) Keys() []string {
	count := m.Count()
	ch := make(chan string, count)

	go func() {
		wg := sync.WaitGroup{}
		wg.Add(SHARD_COUNT)
		for _, shard := range m {
			go func(shard *ConcurrentMapShared) {
				defer wg.Done()

				shard.RLock()

				for key := range shard.items {
					ch <- key
				}
				shard.RUnlock()
			}(shard)
		}
		wg.Wait()
		close(ch)
	}()
	keys := make([]string, count)
	// 统计各个协程并发读取 Map 分片的 key
	for k := range ch {
		keys = append(keys, k)
	}
	return keys
}
