package main

import (
	"log"
	"time"
)

type FuncType func(args ...interface{}) error

// 模拟接口请求发起
func doLimiterApiRequest(args ...interface{}) error {
	time.Sleep(time.Millisecond * 1000)
	log.Println("api request")
	log.Println(args...)
	return nil
}

/*
 * 并发限速装饰器
 */
func limiterApiDecorator(apiFunc FuncType, concurrentMaxNum int, args ...interface{}) {
	limiter := make(chan int, concurrentMaxNum)
	for {
		limiter <- 1
		go func() {
			_ = apiFunc(args...)
			<-limiter
		}()
	}
}

func main() {
	limiterApiDecorator(doLimiterApiRequest, 3)
	limiterApiDecorator(doLimiterApiRequest, 3, 1, 2, 3)
}
