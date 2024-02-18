package main

import (
	"fmt"
	"time"
)

var x int64 = 0x3333333333333333

func storeFunc() {
	for i := 0; ; i++ {
		if i%2 == 0 {
			x = 0x1111111111111111
		} else {
			x = 0x2222222222222222
		}
		// time.Sleep(time.Second)
	}
}

func main() {
	go storeFunc()

	// 对x的赋值并没有写回到内容(可能就存储在寄存器中)，而storeFunc和main是在两个goroutine里面运行的，
	// 他们并不共享CPU执行上下文，从而main读出的值永远是初始值。
	// 如果在storeFunc的循环内加一个sleep间隔，让x的值能够被写入内容，那么以下会得到期望的值
	for {
		time.Sleep(time.Second)
		fmt.Printf("%x\n", x) // 始终是：0x3333333333333333
	}
}
