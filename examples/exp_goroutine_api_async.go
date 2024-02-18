package main

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

/*
 * 接口异步接口处理
 * 业务场景：接口本身没有超时，但是会对期望结果进行超时处理
 */

const (
	StatusPending = "pending" // 准备
	StatusProcess = "process" // 进行
	StatusSuccess = "success" // 成功
	StatusFailure = "failure" // 失败
)

type RsyncApiData string // 可根据实际接口重新配置
type RsyncApiError error

type DecoratorRsyncApiResult struct {
	ApiData  RsyncApiData
	ApiError RsyncApiError
}

// 模拟接口请求发起
func doRsyncApiRequest() (RsyncApiData, RsyncApiError) {
	return StatusPending, nil
}

// 模拟接口异步结果（pending）
func doRsyncApiResponsePending() (RsyncApiData, RsyncApiError) {
	return StatusPending, nil
}

// 模拟接口异步结果（process）
func doRsyncApiResponseProcess() (RsyncApiData, RsyncApiError) {
	return StatusProcess, nil
}

// 模拟接口异步结果（success）
func doRsyncApiResponseSuccess() (RsyncApiData, RsyncApiError) {
	return StatusSuccess, nil
}

// 模拟接口异步结果（failure）
func doRsyncApiResponseFailure() (RsyncApiData, RsyncApiError) {
	return StatusFailure, nil
}

// 模拟接口异步结果（异常场景）
func doRsyncApiResponseException() (RsyncApiData, RsyncApiError) {
	return "", errors.New("api exception")
}

// 接口异步结果处理装饰器
func rsyncApiResponseDecorator(apiFunc func() (RsyncApiData, RsyncApiError), expectedDataSlice []RsyncApiData, decoratorResultChan chan DecoratorRsyncApiResult, timeout int) {
	d := time.Duration(timeout) * time.Second
	timeAfter := time.After(d)

	done := make(chan bool) // 请求完成
	for {
		fmt.Println("[response] start")
		go func(decoratorResultChan chan DecoratorRsyncApiResult) {
			res, err := apiFunc()
			fmt.Println("[response] end")
			// 错误处理
			if err != nil {
				decoratorResultChan <- DecoratorRsyncApiResult{res, err}
				done <- true
				return
			}
			// 期望结果
			for _, expectedData := range expectedDataSlice {
				if res == expectedData {
					decoratorResultChan <- DecoratorRsyncApiResult{res, err}
					done <- true
					return
				}
			}
		}(decoratorResultChan)

		select {
		case <-done:
			fmt.Println("[response] get response!")
			return
		case <-timeAfter:
			fmt.Println("[response] timeout...")
			decoratorResultChan <- DecoratorRsyncApiResult{"", errors.New("response timeout")}
			return
		default:
			fmt.Println("[response] sleep")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 发起请求
	_, eReq := doRsyncApiRequest()
	if eReq != nil {
		fmt.Println(eReq.Error())
		return
	}
	// 处理结果
	expectedStatus := []RsyncApiData{StatusSuccess, StatusFailure} // 期望结果
	timeoutResponse := 5                                           // 响应超时时间
	resultChan := make(chan DecoratorRsyncApiResult, 1)

	go rsyncApiResponseDecorator(doRsyncApiResponsePending, expectedStatus, resultChan, timeoutResponse)
	//go rsyncApiResponseDecorator(doRsyncApiResponseProcess, expectedStatus, resultChan, timeoutResponse)
	//go rsyncApiResponseDecorator(doRsyncApiResponseSuccess, expectedStatus, resultChan, timeoutResponse)
	//go rsyncApiResponseDecorator(doRsyncApiResponseFailure, expectedStatus, resultChan, timeoutResponse)
	//go rsyncApiResponseDecorator(doRsyncApiResponsePending, expectedStatus, resultChan, timeoutResponse)
	//go rsyncApiResponseDecorator(doRsyncApiResponseException, expectedStatus, resultChan, timeoutResponse)

	select {
	case decoratorResult := <-resultChan:
		fmt.Println("[response]", "data:", decoratorResult.ApiData, "; error:", decoratorResult.ApiError)
		return
	}
}
