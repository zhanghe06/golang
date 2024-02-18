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

	// 过期时间测试
	keyToken := "test:token"
	fmt.Println(keyToken)
	expire := 3600
	_, err = redisClient.SetNX(ctx, keyToken, "abc", time.Second * time.Duration(expire)).Result()
	if err != nil {
		fmt.Printf("set token failed: %s", err)
		return
	}

	var expireRedis time.Duration
	expireRedis, err = redisClient.TTL(ctx, keyToken).Result()
	if err != nil && err != redis.Nil {
		fmt.Printf("Cache get token TTL failed: %s", err)
		return
	}
	if expireRedis > 0 {
		tokenExpire := time.Now().Unix() + expireRedis.Microseconds() / 1000
		fmt.Println(tokenExpire)
		return
	}

}
