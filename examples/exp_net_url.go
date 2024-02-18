package main

import (
	"fmt"
	"net/url"
)

func main()  {
	testUrl := "https://www.baidu.com/s?wd=golang&name=test.PDF"
	fmt.Println(url.PathEscape(testUrl))
	u, err := url.Parse(testUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Host: ", u.Host)
	fmt.Println("Path: ", u.Path)
	fmt.Println("Query: ", u.Query().Encode())
	fmt.Println("Query: ", u.RawQuery)
}