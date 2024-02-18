package main

import(
	"fmt"
	"net/url"
)

//https://www.w3schools.com/tags/ref_urlencode.ASP

func testQueryEscape() {
	// QueryEscape: 空格会转为 +
	var urlStr string = "中文 空格=.doc"
	escapeUrl := url.QueryEscape(urlStr)
	fmt.Println("编码:",escapeUrl)

	unEscapeUrl, _ := url.QueryUnescape(escapeUrl)
	fmt.Println("解码:",unEscapeUrl)
}

func testPathEscape() {
	// PathEscape: 空格会转为 %20
	var urlStr string = "中文 空格=.doc"
	escapeUrl := url.PathEscape(urlStr)
	fmt.Println("编码:",escapeUrl)

	//unEscapeUrl, _ := url.PathUnescape(escapeUrl)
	unEscapeUrl, _ := url.PathUnescape("%25E4%25B8%25AD%25E6%2596%2587%20%25E7%25A9%25BA%25E6%25A0%25BC.doc")
	fmt.Println("解码:",unEscapeUrl)
}

func main()  {
	//testQueryEscape()
	//testPathEscape()
	// url.Values.Encode() 内部使用的就是 QueryEscape() 方法，所以最终解决方案是在 url.Encode() 后对 + 替换为 %20:
	// strings.Replace(url.Values.Encode(), "+", "%20", -1)

	//escapeUrl := url.PathEscape("!.txt")
	//escapeUrl := url.PathEscape("!=.txt")
	//fmt.Println("编码:",escapeUrl)

	var urlStr string = "%.doc"
	escapeUrl := url.PathEscape(urlStr)
	fmt.Println("编码:",escapeUrl)

	unEscapeUrl, _ := url.PathUnescape(escapeUrl)
	fmt.Println("解码:",unEscapeUrl)
}
