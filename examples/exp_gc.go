package main

import (
	"time"
)

func main() {
	ch := make(chan int, 10)

	go func() {
		var i = 1
		for {
			i++
			ch <- i
		}
	}()

	timeAfter := time.After(3 * time.Minute)

	for {
		select {
		case x := <-ch:
			println(x)
		//case <- time.After(3 * time.Minute):
		// 定时器在循环外层定义，避免内存泄露
		case <-timeAfter:
			println(time.Now().Unix())
		}
	}
}
