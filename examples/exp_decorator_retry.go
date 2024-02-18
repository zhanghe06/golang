package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"time"
)

type DecoratorRetryApiResult struct {
	ApiData  string
	ApiError error
}

// 模拟接口请求发起
func doErrorApiRequest() (string, error) {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	numRand := seededRand.Intn(3)
	// 随机返回异常
	switch numRand {
	// 错误
	case 0:
		return "", errors.New("api error")
	// 超时
	case 1:
		time.Sleep(10 * time.Second)
		return "", errors.New("api timeout")
	// 恐慌
	case 2:
		panic("api panic")
	// 正常
	default:
		return "1", nil
	}
}

/*
 * 接口重试装饰器
 * 适用于幂等接口
 */
func retryApiDecorator(apiFunc func() (string, error), decoratorResultChan chan DecoratorRetryApiResult, countRetry, timeout int) {
	done := make(chan bool) // 接口请求完成

	for i := 0; i < countRetry+1; i++ {
		// 显示重试次数
		if i > 0 {
			log.Println()
			log.Println("[request ] request retry:", i)
		}

		// 设置单次请求超时
		d := time.Duration(timeout) * time.Second
		timeAfter := time.After(d)

		// 单次接口结果（必须设置缓冲，防止apiFunc的goroutine阻塞）
		apiResChan := make(chan DecoratorRetryApiResult, 1)
		// 单次接口错误
		apiErrorChan := make(chan bool)

		log.Println("[request ] start")
		go func(apiResChan chan DecoratorRetryApiResult) {
			// panic recover
			defer func() {
				if rec := recover(); rec != nil {
					err := fmt.Errorf("%v", rec)
					apiResChan <- DecoratorRetryApiResult{"", err}
					apiErrorChan <- true
					return
				}
			}()
			res, err := apiFunc()
			apiResChan <- DecoratorRetryApiResult{res, err}

			if err == nil {
				done <- true
				return
			} else {
				apiErrorChan <- true
			}
		}(apiResChan)
		log.Println("[request ] end")

		select {
		// 正常请求结束
		case <-done:
			decoratorResultChan <- <-apiResChan
			log.Println("[response] done!")
			return
		// 单次请求错误
		case <-apiErrorChan:
			log.Println("[response] error")
			errData := <-apiResChan
			log.Println("[response]", "data:", errData.ApiData, "; error:", errData.ApiError)
		// 单次请求超时
		case <-timeAfter:
			log.Println("[response] timeout")
		}
		// 比较已经重试次数（是否达到重试最大次数）
		if i >= countRetry {
			decoratorResultChan <- DecoratorRetryApiResult{"", errors.New("response max retries")}
			return
		}
	}
}

func main() {
	//_, _ = doErrorApiRequest()

	countRetry := 3      // 重试次数（总请求数 = 重试次数 + 1）
	timeoutResponse := 5 // 响应超时时间（单次请求）
	resultChan := make(chan DecoratorRetryApiResult)

	go retryApiDecorator(doErrorApiRequest, resultChan, countRetry, timeoutResponse)

	decoratorResult := <-resultChan
	log.Println("[response]", "data:", decoratorResult.ApiData, "; error:", decoratorResult.ApiError)
}
