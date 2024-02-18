package main

import (
	"context"
	"fmt"
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
		return
	}

	key := "test:set"

	// 清空集合
	_ = redisClient.Del(ctx, key).Err()

	// 添加集合
	var prizeSlice []interface{}
	var i uint64
	i = 1
	for {
		if i > 10 {
			break
		}
		prizeSlice = append(prizeSlice, i)
		i++
	}
	err = redisClient.SAdd(ctx, key, prizeSlice...).Err()

	// 弹出集合
	var val uint64
	val, err = redisClient.SPop(ctx, key).Uint64()
	// Redis错误处理
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	}
	// 2个奖品
	if val == 6 || val == 8 {
		fmt.Printf("win: %d", val)
	} else {
		fmt.Printf("lost: %d", val)
	}
}
