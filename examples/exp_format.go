package main

import (
	"errors"
	"fmt"
)

func main()  {
	// 右对齐，左侧填充空格
	// Content-Type: application/json; charset=utf-8
	// Content-Length: 9
	formatA := "%20s: %v\n"
	fmt.Printf(formatA, "Content-Type", "application/json; charset=utf-8")
	fmt.Printf(formatA, "Content-Length", 9)

	// 左对齐，右侧填充空格
	// Content-Type        : application/json; charset=utf-8
	// Content-Length      : 9
	formatB := "%-20s: %v\n"
	fmt.Printf(formatB, "Content-Type", "application/json; charset=utf-8")
	fmt.Printf(formatB, "Content-Length", 9)


	// 打印字节
	var a []byte
	fmt.Printf("%s\n", a) // 空
	a = []byte("abc")
	fmt.Printf("%s\n", a) // abc
	fmt.Printf("%x\n", a) // 616263
	fmt.Printf("%v\n", a) // [97 98 99]

	// 打印错误
	var e error
	fmt.Printf("%s\n", e) // %!s(<nil>)
	fmt.Printf("%v\n", e) // <nil>
	e = errors.New("this is error")
	fmt.Printf("%s\n", e) // this is error
	fmt.Printf("%v\n", e) // this is error
}
