package main

import "fmt"

func main() {
	var code string

	// 十六进制（默认长度）
	code = fmt.Sprintf("%x", 196)
	fmt.Println(code)

	// 十六进制（指定长度）
	code = fmt.Sprintf("%06x", 196)
	fmt.Println(code)
}
