package main

import (
	"fmt"
	"sync"
	"time"
)

// Go语言标准库中的sync.Once可以保证go程序在运行期间的某段代码只会执行一次，作用与init类似，但是也有所不同:
// init函数是在文件包首次被加载的时候执行，且只执行一次
// sync.Once是在代码运行中需要的时候执行，且只执行一次

// 简单应用场景
// 一.延迟初始化 在调用函数时先进行一次初始化操作再进行执行。
// 二.捕获error 并发执行处理函数，需要捕获第一次出现的错误。

var once sync.Once

func onces() {
	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}

func main() {

	for i, _ := range make([]string, 10) {
		once.Do(onces) // 执行1次
		fmt.Println("count:", i+1)
	}
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onced) // 不会执行
			fmt.Println("---")
		}()
	}
	time.Sleep(time.Second)
}
