package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	n := runtime.GOMAXPROCS(1) // 设定并行计算CPU最大核数，并返回之前的值
	fmt.Printf("n = %d\n", n)  // n = 8
	wg := sync.WaitGroup{}
	// Add(n) n为总共要等待的协程数
	wg.Add(20)
	// A输出是10个10
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done() // 相当于Add(-1)
		}()
	}
	// B输出0-9，但是顺序分情况
	// 不设置runtime.GOMAXPROCS(1)时，B每次输出都会出现乱序
	// 设置runtime.GOMAXPROCS(1)时，在早期Go版本中，B顺序输出0-9
	// 设置runtime.GOMAXPROCS(1)时，在新版Go版本中，B先输出9，然后输出A，最后B顺序输出0-8
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done() // 相当于Add(-1)
		}(i)
	}
	fmt.Println("---main end loop---")
	wg.Wait() // 等待阻塞结束，在1个操作系统线程情况下，此时子协程才开始按入队顺序执行
	fmt.Println("---  main exit  ---")
}

/*
n = 8
---main end loop---
B:  9
A:  10
A:  10
A:  10
A:  10
A:  10
A:  10
A:  10
A:  10
A:  10
A:  10
B:  0
B:  1
B:  2
B:  3
B:  4
B:  5
B:  6
B:  7
B:  8
---  main exit  ---
*/

/*
新版实现中：

每增加一个子协程就把其对应的函数地址存放到”_p_.runnext“，
而把”_p_.runnext“原来的地址（即上一个子协程对应的函数地址）移动到队列”_p_.runq“里面，
这样当执行到wg.Wait()时，”_p_.runnext“存放的就是最后一个子协程对应的函数地址（即输出B: ９的那个子协程）。

当开始执行子协程对应的函数时，首先执行”_p_.runnext“对应的函数，
然后按先进先出的顺序执行队列”_p_.runq“里的函数。所以这就解释了为什么总是B：9打在第一个，而后面打印的则是进入队列的顺序。

参考: https://zhuanlan.zhihu.com/p/258759046
*/
