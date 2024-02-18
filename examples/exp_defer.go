package main

import (
	"fmt"
)

func deferCall() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(1, err) //err 就是panic传入的参数
		}
		fmt.Println("打印前")
	}()

	defer func() { // 必须要先声明defer，否则recover()不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(2, err) //err 就是panic传入的参数
		}
		fmt.Println("打印中")
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(3, err) //err 就是panic传入的参数
		}
		fmt.Println("打印后")
	}()

	panic("触发异常")
}

func main() {
	// 3 触发异常
	// 打印后
	// 打印中
	// 打印前
	deferCall()
}
