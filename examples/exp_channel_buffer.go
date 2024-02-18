package main

import (
	"fmt"
	"sync"
)

func generator(n int) <-chan int {
	outCh := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			outCh <- i
		}
		close(outCh)
	}()
	return outCh
}

func do(inCh <-chan int, outCh chan<- int, wg *sync.WaitGroup) {
	for v := range inCh {
		outCh <- v * v
	}

	wg.Done()
}

func main() {
	inCh := generator(100)
	outCh := make(chan int, 10)

	// 使用5个`do`协程同时处理输入数据
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go do(inCh, outCh, &wg)
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()

	for r := range outCh {
		fmt.Println(r)
	}
}
