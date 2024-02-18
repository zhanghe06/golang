package main

import (
	"fmt"
	"time"
)

// 定时器 ticker
func main() {
	d := time.Duration(time.Second * 2)
	t := time.NewTicker(d)
	defer t.Stop()

	for {
		<-t.C
		time.Sleep(time.Second * 1)
		fmt.Println("timeout...", time.Now().Format("2006-01-02 15:04:05"))
	}
}

// 定时器 timer
/*
func main() {
	d := time.Duration(time.Second*2)
	t := time.NewTimer(d)
	defer t.Stop()

	for {
		<- t.C

		fmt.Println("timeout...")
		// need reset
		t.Reset(time.Second*2)
	}
}
*/
