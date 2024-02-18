package main

import (
	"fmt"
)

func main() {
	// 一、非缓冲channel关闭
	ch := make(chan string)
	close(ch)

	// close(ch) // 会panic的，channel不能close两次

	// ch <- "g" // 会panic的，不能向已经关闭的channel写数据

	// 无缓冲channel，读取被关闭的channel
	val := <- ch // 不会panic, i读取到的值是空 "",  如果channel是bool的，那么读取到的是false
	fmt.Println("val:", val) // val:

	// 二、缓冲channel关闭
	cb := make(chan string, 5)
	cb <- "test"
	close(cb)

	val = <- cb // 缓冲channel未读取完毕，可以继续读取channel中的剩余的数据
	fmt.Println("val:", val) // val: test

	val = <- cb // 缓冲channel已读取完毕，读取到的值是空 ""
	fmt.Println("val:", val) // val:
}


