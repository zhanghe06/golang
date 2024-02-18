package main

import (
	"context"
	"fmt"
	"time"
)

/*
 * 创建一个管道chan，启动goroutine
 * for循环存数据
 */
func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				// 执行defer cancel操作后，就会执行到该select入库
				fmt.Println("i exited")
				return // returning not to leak the goroutine
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func testContextWithCancel() {

	ctx, cancel := context.WithCancel(context.Background())
	// 当取数据n == 5时候，执行defer cancel操作
	defer cancel()
	intChan := gen(ctx)
	for n := range intChan {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func main() {
	testContextWithCancel()
	time.Sleep(time.Second)
}
