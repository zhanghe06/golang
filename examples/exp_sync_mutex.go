package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	// 读写操作都需要加锁，如果读操作未加锁，并发读写时，可能会出现报错状况:
	// 当并发数量超过20时，会panic: fatal error: concurrent map read and map write
	ua.Lock()
	defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func testMutexMap() {
	// n太小时不会（比如20以内），因机器而异
	// fatal error: concurrent map read and map write
	n := 200
	//m := make(map[string]int)
	ua := &UserAges{
		ages: map[string]int{},
	}
	//forever := make(chan bool)

	go func() {
		for i := 0; i < n; i++ {
			fmt.Println(i, ua.Get(strconv.Itoa(i))) // read
		}
	}()

	go func() {
		for i := 0; i < n; i++ {
			ua.Add(strconv.Itoa(i), i) // write
		}
	}()

	time.Sleep(time.Second * 5)
	//<-forever
}

func main() {
	testMutexMap()
}
