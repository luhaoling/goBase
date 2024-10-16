package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

var (
	redisClient *redis.Client
	Producer    = &kafka.Writer{}
)

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   5,
	})
	fmt.Println("Redis client initialized")

	Producer = &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "test",
		Balancer:               &kafka.Hash{},
		MaxAttempts:            0,
		WriteBackoffMin:        0,
		WriteBackoffMax:        0,
		BatchSize:              0,
		BatchBytes:             0,
		BatchTimeout:           0,
		ReadTimeout:            0,
		WriteTimeout:           time.Second,
		RequiredAcks:           kafka.RequireNone,
		Async:                  false,
		Completion:             nil,
		Compression:            0,
		Logger:                 nil,
		ErrorLogger:            nil,
		Transport:              nil,
		AllowAutoTopicCreation: false,
	}
}

func acquireLock(lockKey string) bool {
	result, err := redisClient.SetNX(context.Background(), lockKey, "locked", 10*time.Second).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("lock result:", result, "lock key:")
	return result
}

func releaseLock(lockKey string) {
	err := redisClient.Del(context.Background(), lockKey).Err()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	defer func() {
		if err := Producer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	lockKey := "my_lock_key"

	if acquireLock(lockKey) {
		fmt.Println("Lock acquired")
		defer releaseLock(lockKey)

		msg := kafka.Message{
			Topic:         "",
			Partition:     0,
			Offset:        0,
			HighWaterMark: 0,
			Key:           []byte(fmt.Sprintf("%s", lockKey)),
			Value:         []byte(fmt.Sprintf("Lock acquired for key: %s", lockKey)),
			Headers:       nil,
			WriterData:    nil,
			Time:          time.Time{},
		}

		err := Producer.WriteMessages(context.Background(), msg)
		fmt.Println("Message sent")
		if err != nil {
			log.Fatal()
		}

		fmt.Println("Lock acquired Performing critical section operations...")
		time.Sleep(5 * time.Second)
		fmt.Println("Critical section operations completed.")
	} else {
		fmt.Println("Failed to acquire lock. Another process may be holding the lock.")
	}

}
