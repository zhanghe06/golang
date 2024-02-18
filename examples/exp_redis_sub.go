package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
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
	ctx := context.Background()
	//ctx, cancel := context.WithTimeout(ctxRoot, 5*time.Second)
	//defer cancel()
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

	for {
		subscriber := redisClient.Subscribe(ctx, channel)
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			log.Println(err)
			// 断线重连流程优化
			// _ = subscriber.Close()
			time.Sleep(redisClient.Options().DialTimeout) // 默认5秒
			continue
		}

		message := Message{}

		if err := json.Unmarshal([]byte(msg.Payload), &message); err != nil {
			log.Println(err)
			continue
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Printf("%+v\n", message)
	}
}
