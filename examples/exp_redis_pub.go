package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisClient https://github.com/go-redis/redis
// 注意：import时别忘了v8，v8支持context，对应Redis 6版本

func main() {
	// 单节点
	redisClient := redis.NewClient(&redis.Options{
		Addr:       "0.0.0.0:6379",
		Password:   "iu$~yaw=:CqX",
		DB:         0,
		MaxRetries: 2,  // 失败请求1+重试请求2=总共请求3
		PoolSize:   10, // 连接池大小
	})
	ctxRoot := context.Background()
	ctx, cancel := context.WithTimeout(ctxRoot, 5*time.Second)
	defer cancel()
	// 检测连接
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
		return
	}

	channel := "room:01"

	type Message struct {
		Name string `json:"name"`
	}

	message := Message{
		Name: "A",
	}

	payload, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	redisClient.Publish(ctx, channel, payload)
}
