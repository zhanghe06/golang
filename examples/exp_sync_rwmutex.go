package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type concurrentMap struct {
	m map[string]int
	//lock *sync.Mutex
	lock *sync.RWMutex
}

func (c *concurrentMap) Get(key string) int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.m[key]
}

func (c *concurrentMap) Set(key string, val int) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.m[key] = val
}

func testMapRWMutex() {
	n := 10
	m := make(map[string]int)
	//lock := new(sync.Mutex)
	lock := new(sync.RWMutex)
	cm := concurrentMap{
		m:    m,
		lock: lock,
	}

	go func() {
		for i := 0; i < n; i++ {
			cm.Set(strconv.Itoa(i), i)
		}
	}()

	go func() {
		for i := 0; i < n; i++ {
			fmt.Println(i, cm.Get(strconv.Itoa(i)))
		}
	}()

	time.Sleep(time.Second * 5)
}

func main() {
	testMapRWMutex()
}
