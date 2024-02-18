package main

import (
	"fmt"
	"time"
)

/**
select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。
select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。
*/

func main() {
	timeout := make(chan bool, 1)
	ch := make(chan int, 1)

	go func() {
		num := time.Now().Second()
		randBool := num%2 == 0
		if randBool {
			ch <- num
			close(ch)
		}
		time.Sleep(1e9) // sleep one second
		timeout <- true // true/false对channel来说是一样的，只是true是有意义的值
	}()

	var i int
	select {
	case i = <-ch:
		fmt.Println(i)
	//case ch <- 100: // 这种写法常用场景：任务队列长度限制
	//	fmt.Println(<-ch)
	case <-timeout:
		fmt.Println("timeout!")
	}
}

/*
比如我们有一个服务， 当请求进来的时候我们会生成一个 job 扔进 channel， 由其他协程从 channel 中获取 job 去执行。
但是我们希望当 channel 满了的时候， 将该 job 抛弃并回复 【服务繁忙，请稍微再试。】 就可以用 select 实现该需求。
*/
