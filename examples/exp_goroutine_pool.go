package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker 任务类型接口
type Worker interface {
	Task(goId int)
}

// Pool 任务池
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New 新建
func New(maxGoroutines int) *Pool {
	//任务池
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)
	//创建maxGoroutines个go协程
	for i := 0; i < maxGoroutines; i++ {
		go func(goId int) {
			//保证goroutine不停止执行通道中的任务
			for w := range p.work {
				w.Task(goId)
			}
			//每个goroutine不再执行work通道中任务时停止
			p.wg.Done()
		}(i)
	}
	return &p
}

// Run 运行
func (p *Pool) Run(r Worker) {
	p.work <- r
}

// Shutdown 停止
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}

var names = []string{
	"lili",
	"lucy",
}

// Worker实现类型
type namePrinter struct {
	name string
}

func (n *namePrinter) Task(goId int) {
	fmt.Printf("goroutineID:%d，打印名字为：%s\n", goId, n.name)
	time.Sleep(time.Second)
}

// 判断接口是否已经实现
var _ Worker = &namePrinter{}

func main() {
	p := New(3)
	var wg sync.WaitGroup
	wg.Add(10 * len(names))

	for i := 0; i < 10; i++ {
		for _, name := range names {
			//任务实例
			np := namePrinter{
				name: name,
			}

			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	p.Shutdown()
}
