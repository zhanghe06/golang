package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"io/ioutil"
	"net/http"
	netUrl "net/url"
	"runtime"
	"time"
)

//type LogInfo struct {
//
//}


func stack() []byte {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			return buf[:n]
		}
		buf = make([]byte, 2*len(buf))
	}
}

func apiGet(ctx context.Context, url string, headers map[string]string) (statusCode int, resBody []byte, err error) {

	defer func() {
		level := "info"
		resStatusCode := ctx.Value("res_status_code")
		if resStatusCode == nil {
			resStatusCode = 0
		}
		resErrorMsg := ctx.Value("res_error_msg")
		if resErrorMsg == nil {
			resErrorMsg = ""
		}
		resErrorStack := ""
		if resStatusCode.(int) < 200 || resStatusCode.(int) > 300 || resErrorMsg != "" {
			level = "error"
		}
		if rec := recover(); rec != nil {
			// 获取全部堆栈
			level = "panic"
			resErrorMsg = fmt.Errorf("%v", rec).Error()
			resErrorStack = string(stack())
		}
		timeEnd := time.Now()             // 结束时间

		requestId := ctx.Value("req_id")
		if requestId == nil || requestId == "" {
			requestId = uuid.Must(uuid.NewV4()).String()
		}

		timeStart := ctx.Value("time_start").(time.Time)
		latency := timeEnd.Sub(timeStart) // 执行时间

		reqHost := ctx.Value("req_host")
		reqMethod := ctx.Value("req_method")
		reqPath := ctx.Value("req_path")
		reqQuery := ctx.Value("req_query")
		reqHeader := ctx.Value("req_header")
		reqBody := ctx.Value("req_body")

		resBodyString := ctx.Value("res_body")
		clientIp := ctx.Value("client_ip")

		logInfo := make(map[string]interface{})
		logInfo["req_id"] = requestId
		logInfo["service_name"] = "go project"
		logInfo["time_start"] = timeStart
		logInfo["time_end"] = timeEnd
		logInfo["latency"] = latency
		logInfo["level"] = level
		logInfo["client_ip"] = clientIp
		logInfo["req_host"] = reqHost
		logInfo["req_method"] = reqMethod
		logInfo["req_path"] = reqPath
		logInfo["req_query"] = reqQuery
		logInfo["req_header"] = reqHeader
		logInfo["req_body"] = reqBody
		logInfo["res_status_code"] = resStatusCode
		logInfo["res_error_msg"] = resErrorMsg
		logInfo["res_error_stack"] = resErrorStack
		logInfo["res_body"] = resBodyString

		logInfoByte, _ := json.Marshal(logInfo)
		fmt.Println(string(logInfoByte))
	}()

	timeStart := time.Now() // 开始时间
	ctx = context.WithValue(ctx, "time_start", timeStart)

	u, err := netUrl.Parse(url)
	if err != nil {
		ctx = context.WithValue(ctx, "res_error_msg", err.Error())
		return
	}
	ctx = context.WithValue(ctx, "req_method", "GET")
	ctx = context.WithValue(ctx, "req_host", u.Host)
	ctx = context.WithValue(ctx, "req_path", u.Path)
	ctx = context.WithValue(ctx, "req_query", u.RawQuery)
	ctx = context.WithValue(ctx, "req_body", "")

	// 创建请求
	req, _ := http.NewRequest("GET", url, nil)
	// 设置头部
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	reqHeaderBytes, err := json.Marshal(req.Header)
	if err != nil {
		ctx = context.WithValue(ctx, "res_error_msg", err.Error())
		return
	}
	reqHeader := string(reqHeaderBytes)
	ctx = context.WithValue(ctx, "req_header", reqHeader)
	// 执行请求
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		ctx = context.WithValue(ctx, "res_error_msg", err.Error())
		return
	}
	defer func() {
		_ = res.Body.Close()
	}()
	// 读取响应
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		ctx = context.WithValue(ctx, "res_error_msg", err.Error())
		return
	}
	statusCode = res.StatusCode

	ctx = context.WithValue(ctx, "res_status_code", statusCode)
	ctx = context.WithValue(ctx, "res_body", string(resBody))

	return
}


func apiPost(ctx context.Context, url string, headers map[string]string, reqBody []byte) (statusCode int, resBody []byte, err error) {
	defer func() {
		level := "info"
		resStatusCode := ctx.Value("res_status_code")
		if resStatusCode == nil {
			resStatusCode = 0
		}
		resErrorMsg := ctx.Value("res_error_msg")
		if resErrorMsg == nil {
			resErrorMsg = ""
		}
		resErrorStack := ""
		if resStatusCode.(int) < 200 || resStatusCode.(int) > 300 || resErrorMsg != "" {
			level = "error"
		}
		if rec := recover(); rec != nil {
			// 获取全部堆栈
			level = "panic"
			resErrorMsg = fmt.Errorf("%v", rec).Error()
			resErrorStack = string(stack())
		}
		timeEnd := time.Now()             // 结束时间

		requestId := ctx.Value("req_id")
		if requestId == nil || requestId == "" {
			requestId = uuid.Must(uuid.NewV4()).String()
		}

		timeStart := ctx.Value("time_start").(time.Time)
		latency := timeEnd.Sub(timeStart) // 执行时间

		reqHost := ctx.Value("req_host")
		reqMethod := ctx.Value("req_method")
		reqPath := ctx.Value("req_path")
		reqQuery := ctx.Value("req_query")
		reqHeader := ctx.Value("req_header")
		reqBodyString := ctx.Value("req_body")

		resBodyString := ctx.Value("res_body")
		clientIp := ctx.Value("client_ip")

		logInfo := make(map[string]interface{})
		logInfo["req_id"] = requestId
		logInfo["service_name"] = "go project"
		logInfo["time_start"] = timeStart
		logInfo["time_end"] = timeEnd
		logInfo["latency"] = latency
		logInfo["level"] = level
		logInfo["client_ip"] = clientIp
		logInfo["req_host"] = reqHost
		logInfo["req_method"] = reqMethod
		logInfo["req_path"] = reqPath
		logInfo["req_query"] = reqQuery
		logInfo["req_header"] = reqHeader
		logInfo["req_body"] = reqBodyString
		logInfo["res_status_code"] = resStatusCode
		logInfo["res_error_msg"] = resErrorMsg
		logInfo["res_error_stack"] = resErrorStack
		logInfo["res_body"] = resBodyString

		logInfoByte, _ := json.Marshal(logInfo)
		fmt.Println(string(logInfoByte))
	}()

	timeStart := time.Now() // 开始时间
	ctx = context.WithValue(ctx, "time_start", timeStart)

	u, err := netUrl.Parse(url)
	if err != nil {
		ctx = context.WithValue(ctx, "res_error_msg", err.Error())
		return
	}
	ctx = context.WithValue(ctx, "req_method", "POST")
	ctx = context.WithValue(ctx, "req_host", u.Host)
	ctx = context.WithValue(ctx, "req_path", u.Path)
	ctx = context.WithValue(ctx, "req_query", u.RawQuery)
	ctx = context.WithValue(ctx, "req_body", string(reqBody))

	// 创建请求
	req, _ := http.NewRequest("POST", url, bytes.NewReader(reqBody))
	// 设置头部
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	reqHeaderBytes, err := json.Marshal(req.Header)
	if err != nil {
		ctx = context.WithValue(ctx, "res_error_msg", err.Error())
		return
	}
	reqHeader := string(reqHeaderBytes)
	ctx = context.WithValue(ctx, "req_header", reqHeader)
	// 执行请求
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		ctx = context.WithValue(ctx, "res_error_msg", err.Error())
		return
	}
	defer func() {
		_ = res.Body.Close()
	}()
	// 读取响应
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		ctx = context.WithValue(ctx, "res_error_msg", err.Error())
		return
	}
	statusCode = res.StatusCode

	ctx = context.WithValue(ctx, "res_status_code", statusCode)
	ctx = context.WithValue(ctx, "res_body", string(resBody))

	return
}


func callGet(ctx context.Context)  {
	url := "https://anyshare.aishu.cn/api/appstore/v1/applist2"
	statusCode, resBody, err := apiGet(ctx, url, nil)
	fmt.Println("statusCode: ", statusCode)
	fmt.Println("resBody: ", string(resBody))
	fmt.Println("error: ", err)
}

func callPost(ctx context.Context)  {
	url := "https://anyshare.aishu.cn/api/appstore/v1/applist2"
	headers := make(map[string]string)
	headers["a"] = "1"
	headers["b"] = "2"

	reqBody := []byte(`{"One":"Two"}`)

	statusCode, resBody, err := apiPost(ctx, url, headers, reqBody)
	fmt.Println("statusCode: ", statusCode)
	fmt.Println("resBody: ", string(resBody))
	fmt.Println("error: ", err)
}

func main() {
	ctx := context.Background()
	//callGet(ctx)
	callPost(ctx)
}
