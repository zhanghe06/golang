package main

import (
	"fmt"
	"time"
)

/*
 * panic 会停掉当前正在执行的程序（注意，不只是协程）。
 * 但是与 os.Exit(-1) 这种直愣愣的退出不同，panic 的撤退比较有秩序。
 * 他会先处理完当前 goroutine 已经 defer 挂上去的任务，执行完毕后再退出整个程序。
 *
 * panic 仅保证当前 goroutine 下的 defer 都会被调到，但不保证其他协程的 defer 也会调到。
 * 以下示例中，main defer 并未执行
 *
 * 结论:
 * 1. goroutine panic 会引发整个程序退出
 * 2. goroutine panic 外部无法直接捕获
 * 3. goroutine panic 可以在内部通过 defer 来 recover
 */
func main() {
	defer func() {
		fmt.Println("main defer") // 执行不到
	}()
	go func() {
		defer func() {
			fmt.Println("goroutine defer")
			// 在通常业务中正确方式：取消以下注释，将 panic 截获，以 error 或字符串形式往上层传递（通过 channel）
			// if rec := recover(); rec != nil {
			// 	err := fmt.Errorf("%v", rec)
			// 	fmt.Println(err.Error())
			// 	return
			// }
		}()

		fmt.Println("goroutine info")
		// goroutine panic
		panic("goroutine panic")
	}()

	fmt.Println("main info")
	time.Sleep(1 * time.Second)
	fmt.Println("main sleep") // 执行不到

	// main panic
	panic("main panic") // 执行不到
}

/*
 * 执行结果(开启sleep)
 * -----------------------------
 * main info
 * goroutine info
 * goroutine defer
 * panic: goroutine panic
 *
 * goroutine 4 [running]:
 * main.main.func1()
 * 	examples/exp_panic.go:20 +0xb2
 * created by main.main
 * 	examples/exp_panic.go:13 +0x39
 */

/*
 * 执行结果(关闭sleep)
 * -----------------------------
 * main info
 * main sleep
 * panic: main panic
 *
 * goroutine 1 [running]:
 * main defer
 * main.main()
 * examples/exp_panic.go:40 +0x127
 */
