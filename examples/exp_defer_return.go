package main

import "fmt"

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

/*
 * http://golang.org/ref/spec#defer_statements
 * return 语句不是一条原子语句，return xxx 其实是赋值 ＋ return 指令
 */
func main() {
	f1 := f1()
	f2 := f2()
	f3 := f3()

	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
}
