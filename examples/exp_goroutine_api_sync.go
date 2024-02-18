package main

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

/*
 * 外部同步接口处理
 * 业务场景：外部接口耗时，需要转为异步处理结果
 */

type SyncApiData string
type SyncApiError error

type DecoratorSyncApiResult struct {
	ApiData  SyncApiData
	ApiError SyncApiError
}

// 模拟外部请求（正常场景）
func doSyncApiSuccess() (SyncApiData, SyncApiError) {
	//time.Sleep(time.Second)
	return "api data", nil
}

// 模拟外部请求（异常场景）
func doSyncApiFailure() (SyncApiData, SyncApiError) {
	time.Sleep(time.Second)
	return "", errors.New("api failure")
}

// 模拟外部请求（超时场景）
func doSyncApiTimeout() (SyncApiData, SyncApiError) {
	time.Sleep(10 * time.Second)
	return "", errors.New("api timeout")
}

// 接口异步请求装饰器
func syncApiRequestDecorator(apiFunc func() (SyncApiData, SyncApiError), decoratorResultChan chan DecoratorSyncApiResult, timeout int) {
	d := time.Duration(timeout) * time.Second
	timeAfter := time.After(d)

	fmt.Println("[request ] start")
	done := make(chan bool) // 外部接口请求完成

	go func(decoratorResultChan chan DecoratorSyncApiResult) {
		res, err := apiFunc()
		decoratorResultChan <- DecoratorSyncApiResult{res, err}
		done <- true
	}(decoratorResultChan)

	fmt.Println("[request ] end")

	select {
	case <-done:
		fmt.Println("[request ] get response!")
		return
	case <-timeAfter:
		fmt.Println("[request ] timeout...")
		decoratorResultChan <- DecoratorSyncApiResult{"", errors.New("timeout")}
		return
	}
}

// 接口异步结果处理器
func syncApiResultHandler(decoratorResultChan chan DecoratorSyncApiResult, done chan bool) {
	for {
		select {
		case decoratorResult := <-decoratorResultChan:
			fmt.Println("[response]", "data:", decoratorResult.ApiData, "; error:", decoratorResult.ApiError)
			done <- true
			return
		default:
			fmt.Println("[response] no message received")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 一、普通方式（阻塞）
	//_, e := doSyncApiFailure()
	//if e != nil{
	//	fmt.Println(e.Error())
	//}

	timeoutRequest := 5
	resultChan := make(chan DecoratorSyncApiResult, 1)

	// 二、同步协程方式（阻塞）
	//go func(decoratorSyncApiResultChan chan DecoratorSyncApiResult) {
	//	res, err := doSyncApiFailure()
	//	decoratorSyncApiResultChan <- DecoratorSyncApiResult{res, err}
	//	close(decoratorSyncApiResultChan)
	//}(resultChan)
	//
	//apiResult := <-resultChan
	//if apiResult.ApiError != nil {
	//	fmt.Println(apiResult.ApiError)
	//}

	// 三、异步协程方式（非阻塞）
	forever := make(chan bool)

	// 外部请求 - 异步方式发起（三种场景：正常、失败、超时）
	//go syncApiRequestDecorator(doSyncApiSuccess, resultChan, timeoutRequest)
	go syncApiRequestDecorator(doSyncApiFailure, resultChan, timeoutRequest)
	//go syncApiRequestDecorator(doSyncApiTimeout, resultChan, timeoutRequest)

	// 外部请求 - 异步处理结果
	go syncApiResultHandler(resultChan, forever)

	fmt.Printf("[*] Waiting for logs. To exit press CTRL+C\n")

	<-forever
}
