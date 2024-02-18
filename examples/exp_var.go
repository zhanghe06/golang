package main

import "fmt"

//v := 1 // 简短声明并赋值，由编译器推导类型，只能用户函数内部 syntax error: non-declaration statement outside function body

//var v int
//v = 1 // syntax error: non-declaration statement outside function body

var v = 1

// 错误写法
// examples/exp_var.go:11:14: undefined: a
//func main()  {
//	if true {
//		a := 1
//	} else {
//		a := 2
//	}
//	fmt.Println(a)
//}

func main() {
	//v = 1 可在外部声明，内部赋值
	fmt.Println(v)

	one := 0
	//one := 1 // error: no new variables on left side of :=
	fmt.Println(one)

	one, two := 1, 2 // two 是新变量，允许 one 的重复声明。比如 error 处理经常用同名变量 err
	fmt.Println(one)
	fmt.Println(two)

	var a int32
	if true {
		a = 1
	} else {
		a = 2
	}
	fmt.Println(a)
}
