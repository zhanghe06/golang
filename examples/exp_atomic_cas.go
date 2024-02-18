package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int32          //计数器
	casWg   sync.WaitGroup //信号量
)

func incCounter(index int) {
	defer casWg.Done()

	// 自旋次数
	spinNum := 0
	// 这里之所以使用无限循环是因为在高并发下每个线程执行CAS并不是每次都成功。
	// 失败了的线程需要重写获取变量当前的值，然后重新执行CAS操作。
	for {
		old := counter
		// CompareAndSwapInt32函数在被调用之后会先判断参数addr指向的被操作值与参数old的值是否相等。
		// 仅当此判断得到肯定的结果之后，该函数才会用参数new代表的新值替换掉原先的旧值。否则，后面的替换操作就会被忽略。
		ok := atomic.CompareAndSwapInt32(&counter, old, old+1)
		if ok {
			break
		} else {
			spinNum++
		}
	}
	fmt.Printf("thread,%d,spinnum,%d\n", index, spinNum)

}

func main() {
	threadNum := 10000

	casWg.Add(threadNum)

	for i := 0; i < threadNum; i++ {
		go incCounter(i)
	}

	casWg.Wait()

	fmt.Println(counter)
}

/*
1、atomic中的操作是原子性的；
2、原子操作由底层硬件支持，而锁则由操作系统的调度器实现。锁应当用来保护一段逻辑，对于一个变量更新的保护，原子操作通常会更有效率，并且更能利用计算机多核的优势，如果要更新的是一个复合对象，则应当使用atomic.Value封装好的实现。
3、atomic中的代码，主要还是依赖汇编来来实现的原子操作。
*/
