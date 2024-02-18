package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

var HttpClient *http.Client


// Transport默认配置
// https://pkg.go.dev/net/http#DefaultTransport

//var DefaultTransport RoundTripper = &Transport{
//	Proxy: ProxyFromEnvironment,
//	DialContext: defaultTransportDialContext(&net.Dialer{
//		Timeout:   30 * time.Second,
//		KeepAlive: 30 * time.Second,
//	}),
//	ForceAttemptHTTP2:     true,
//	MaxIdleConns:          100,
//	IdleConnTimeout:       90 * time.Second,
//	TLSHandshakeTimeout:   10 * time.Second,
//	ExpectContinueTimeout: 1 * time.Second,
//}

func initHTTPClient() {
	HttpClient = new(http.Client)
	tr := &http.Transport{
		// DisableKeepAlives: false, // FIXME: connect reset by peer
		// https://stackoverflow.com/questions/37774624/go-http-get-concurrency-and-connection-reset-by-peer
		MaxIdleConns:       100,
		MaxIdleConnsPerHost:  100, // 默认2
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		IdleConnTimeout: 90 * time.Second, // 它不控制client request的阻塞阶段，但是可以控制连接池中一个连接可以idle多长时间。
	}
	HttpClient.Transport = tr
}


func httpPoolGet(url string)  {
	// 请求超时设置
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	req = req.WithContext(ctx)
	//req.Close = true // FIXME: connect reset by peer
	res, err := HttpClient.Do(req)

	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("code: %d, url: %s\n", res.StatusCode, res.Request.URL)
	//resByte, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(resByte))
}

func main() {
	initHTTPClient()
	i := 1
	for {
		url := "https://www.baidu.com/s?wd=a"
		go httpPoolGet(url)
		i++
		time.Sleep(1 * time.Millisecond)
	}
}

// Tuning the Go HTTP Client Settings for Load Testing
// http://tleyden.github.io/blog/2016/11/21/tuning-the-go-http-client-library-for-load-testing/
