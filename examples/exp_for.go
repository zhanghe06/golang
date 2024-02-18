package main

import (
	"fmt"
	"time"
)

func forWorker(i int) {
	defer func() {
		if rec := recover(); rec != nil {
			err := fmt.Errorf("%v", rec)
			fmt.Printf("%02d: %s\n", i, err.Error())
			return
		}
	}()
	for {
		fmt.Printf("%02d: worker\n", i)
		panic("worker panic")
	}
}

/*
 * 守护 goroutine
 * 捕获 goroutine 的 panic，使其正常退出，然后循环
 */
func main() {
	i := 1
	for {
		go forWorker(i)
		i++
		time.Sleep(1 * time.Second)
	}
}
