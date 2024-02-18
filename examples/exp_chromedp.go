package main

// https://github.com/chromedp/chromedp
// 参考: https://www.bilibili.com/read/cv13239592/
// 参考: https://mojotv.cn/2018/12/26/chromedp-tutorial-for-golang
// 参考: https://studygolang.com/topics/12596

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/cdproto/input"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

var cnt = 0

func main() {
	// 定义打开谷歌浏览器的一个临时数据文件夹
	dir, err := ioutil.TempDir("", "chromedp-example")
	if err != nil {
		panic(err)
	}
	// defer函数代表当整个程序执行完之后会执行os.RemoveAll(dir)，其实就是把这个临时数据文件夹删除
	defer os.RemoveAll(dir)

	// 配置一下等会程序运行打开的浏览器的一些参数
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		// 禁止GPU
		chromedp.DisableGPU,
		// 禁用默认的浏览器检查
		chromedp.NoDefaultBrowserCheck,
		// 一般我们调试的时候肯定是将这个值置为false,这样你就能看到程序在运行时打开浏览器，如果需要部署
		// 到服务器，你希望无头打开浏览器，就得把这个值置为true
		chromedp.Flag("headless", false),
		// 忽略证书错误
		chromedp.Flag("ignore-certificate-errors", true),
		// 使用刚才创建的临时数据文件夹
		chromedp.UserDataDir(dir),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	// 最后执行完之后肯定会关闭这个上下文
	defer cancel()

	// 使用log.Printf打印日志
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// 检查浏览器进程是否启动
	if err := chromedp.Run(taskCtx); err != nil {
		panic(err)
	}

	// 监听网络事件
	listenForNetworkEvent(taskCtx)
	// actions就代表后面打开浏览器要执行的一系列操作
	var actions []chromedp.Action

	actions = append(actions, network.Enable())
	// 指定要访问的地址
	actions = append(actions, chromedp.Navigate(`https://image.baidu.com/search/index?tn=baiduimage&ipn=r&ct=201326592&cl=2&lm=-1&st=-1&fm=result&fr=&sf=1&fmq=1631628760308_R&pv=&ic=&nc=1&z=&hd=&latest=&copyright=&se=1&showtab=0&fb=0&width=&height=&face=0&istype=2&ie=utf-8&sid=&word=%E7%BE%8E%E5%A5%B3%E5%A4%B4%E5%83%8F`))

	// 模拟滚轮滚动50次，触发新的图片加载
	for i := 0; i < 20; i++ {
		actions = append(actions, chromedp.Sleep(1*time.Second))
		actions = append(actions, chromedp.ActionFunc(func(ctx context.Context) error {
			time.Sleep(1 * time.Second)
			// 在页面的（200，200）坐标的位置
			p := input.DispatchMouseEvent(input.MouseWheel, 200, 200)
			p = p.WithDeltaX(0)
			// 滚轮向下滚动1000单位
			p = p.WithDeltaY(float64(1000))
			err = p.Do(ctx)
			return err
		}))
	}

	// 执行这一列的操作
	chromedp.Run(taskCtx,
		actions...,
	)

}

//监听网络事件
func listenForNetworkEvent(ctx context.Context) {
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		// 是一个响应收到的事件
		case *network.EventResponseReceived:
			resp := ev.Response
			if len(resp.Headers) != 0 {
				//将这个resp转成json
				response, _ := resp.MarshalJSON()
				var res = &UrlResponse{}
				json.Unmarshal(response, &res)
				// 我们只关心是图片地址的url
				if strings.Contains(res.Url, ".jpg") || strings.Contains(res.Url, "f=JPEG") {
					cnt++
					// 去对每个图片地址下载图片
					downloadImage(res.Url, "美女头像", cnt)
				}
			}
		}
	})
}

type UrlResponse struct {
	Url string `json:"url"`
}

/**
根据图片的地址下载图片
*/
func downloadImage(imgUrl, dir string, cnt int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("发生异常，地址忽略: %s", imgUrl)
		}
	}()
	//生成文件名
	fileName := path.Base(strconv.Itoa(cnt) + ".jpg")

	// 设置请求地址和请求头参数
	imgReq, err := http.NewRequest("GET", imgUrl, nil)
	imgReq.Header.Add("Referer", "https://image.baidu.com/")
	imgReq.Header.Add("Accept-Encoding", "gzip,deflate,br")
	imgReq.Header.Add("Host", "image.baidu.com")
	imgReq.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36")

	client := &http.Client{}
	// 执行请求
	imgRes, err := client.Do(imgReq)

	if err != nil {
		log.Println("Get image error :", err)
		return
	}
	defer imgRes.Body.Close()
	if imgRes.ContentLength == 0 {
		return
	}

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatalln("Create dir error:", err)
		return
	}

	f, err := os.Create(dir + "/" + fileName)
	if err != nil {
		log.Println("Create image error:", err)
		return
	}
	// 拷贝二进制流数据，保存成本地图片
	io.Copy(f, imgRes.Body)
}
