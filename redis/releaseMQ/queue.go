package queue

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
)

// from: https://learnku.com/articles/87043
// producer 负责发送消息，
// consumer 负责接受和处理消息
// 使用 redis 作为后端存储，利用其高性能和持久化特性来保证消息的可靠性

// 这是一个支持延时消息的消息队列，消息通过 Message.GetScore() 生成排序分数（基于消费时间），存入 redis 有序集合。消费者通过轮询
// 有序集合，仅处理消费时间小于等于当前时间的消息。
// 典型的应用场景：订单超时取消、定时通知等需要延迟触发的任务。

// 如何支持普通消息队列
// 即使消费配置：如果生产者在发布消息时，将 ConsumeTime（消费时间）设置为当前时间或更早的时间，消息会立即被消费者处理。此时的队列类似于普通消息队列
// 局限性：由于消息的调度依赖有序集合的分数（消费时间），无法严格保证先入先出顺序，若多个消息的消费时间相同，排序可能依赖消息 ID，而非入队顺序

// 与普通消息队列的区别
// 消息顺序：
// 普通消息队列：严格按照入队顺序；延迟消息队列：按消费时间顺序
// 核心数据结构
// 普通消息队列：Redis List；延迟消息队列：Redis 有序集合（ZRANGE+哈希表）
// 典型用途：
// 普通消息队列：实时任务处理（如日志异步化）；延迟消息队列：延迟任务调度（如订单取消）

// 为什么需要 zset 和 hash，
// zset 的核心作用：调度延迟消息
// 1. 功能需求：消息队列需要支持延迟消费（例如 30min 后处理订单超时），而 zset 按分数排序的特性天然适合此类场景
// 2. 实现方式：
//	每个消息的消费时间被转换为 zset 的分数（score）
//  消费者通过 ZRABGEBTYSCORE 命令定期查询当前时间之前的所有消息 ID，实现精准调度
// 3. 优势：避免轮询全部消息，仅筛选到期消息，时间复杂度低至 O(log N)

// hash 的核心作用：存储消息详情
//  1. 功能需求：消息包含多个属性（ID、创建时间、消息体等），需要高效存储和快速查询
//  2. 实现方式：
//     每个消息的 ID 作为 hash 的键，消息的完整数据作为值
//     消费者从 zset 获取消息 ID 后，通过 HGET 命令直接提取消息内容
//  3. 优势：hash 的 O(1) 时间复杂查询性能远高于遍历 zset 的原始数据

// 仅使用 zset 实现消息队列的缺点
// 1. 数据冗余和存储效率低
//	每条消息的完整内容直接存储在 Zset 的 member 字段种，导致相同消息在不同延迟时间下重复存储，占用更多内存
// 2. 查询性能下降
//	若消息内容较大，每次通过 zRangeByScore 获取消息时，需传输大量数据，增加网络负载和反序列化开销

// Queue 队列管理器，负责创建队列、发布消息和启动消费者
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
			handler: defaultHandler, // 默认处理函数
		}

		// 应用传入的选项（opts）
		for _, apply := range opts {
			apply(&defaultOptions)
		}

		queue = &Queue{
			ctx:      ctx,
			redis:    redis,
			topic:    defaultOptions.topic,
			producer: NewProducer(ctx),
			consumer: NewConsumer(ctx, defaultOptions.handler), // 创建消费者，使用处理函数
		}

	})

	return queue
}

// Start 启动消费者的监听
func (q *Queue) Start() {
	q.consumer.listen(q.redis, q.topic)
}

// Publish 用于发布消息
func (q *Queue) Publish(msg *Message) (int64, error) {
	return q.producer.publish(q.redis, q.topic, msg)
}
