package main

import (
	"fmt"
	"strconv"
	"time"
)

func mapDemoBuiltIn() {
	defer func() {
		// 注意，这里捕获不到并发读写的goroutine
		// 如果需要捕获，需要通过channel
		if rec := recover(); rec != nil {
			err := fmt.Errorf("%v", rec)
			fmt.Println(err.Error())
		}
	}()
	// n太小时不会（比如20以内），因机器而异
	// fatal error: concurrent map read and map write
	n := 28
	m := make(map[string]int)
	//forever := make(chan bool)

	go func() {
		for i := 0; i < n; i++ {
			m[strconv.Itoa(i)] = i // write
		}
	}()

	go func() {
		for i := 0; i < n; i++ {
			fmt.Println(i, m[strconv.Itoa(i)]) // read
		}
	}()

	time.Sleep(time.Second * 5)
	//<-forever
}

func main() {
	//mapDemoBuiltIn()
	t := make(map[string]int)
	t["a"] = 1
	t["b"] = 2
	fmt.Println(len(t))
}
