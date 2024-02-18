package main

import (
	"flag"
	"fmt"
)

var (
	vl string
	hp string
)

func main() {
	// 参数解析
	flag.StringVar(&vl, "video_link", "", "请输入最后一段视频链接（.ts格式）")
	flag.StringVar(&hp, "http_proxy", "", "请输入HTTP代理")
	flag.Parse()

	fmt.Println(vl)
	fmt.Println(hp)

}
