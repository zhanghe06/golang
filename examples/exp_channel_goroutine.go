package main

import (
	"fmt"
	"time"
)

// channel可以用在goroutine之间的同步
func worker(done chan bool) {
	time.Sleep(time.Second)
	// 通知任务完成
	fmt.Println("通知任务完成")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	// 等待任务完成
	fmt.Println("等待任务完成")
	<-done
}
