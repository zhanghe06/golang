package main

import "fmt"

func main() {
	var p1, p2 int
	fmt.Println(&p1 == &p1) // true 取指针地址，同一变量地址一样
	fmt.Println(&p1 == &p2) // false 取指针地址，两者不一样
	fmt.Println(&p1 == nil) // false
	fmt.Println(p1 == 0)    // true

	// 指针涉及到两个操作符 & 和 * , & 代表 取址符 * 代表 声明符 / 取值符
	// 定义变量a并赋值
	var a int = 1
	// 声明指针变量p 指针类型为int
	var p *int
	// 获取变量a的内存地址赋值给p
	p = &a

	fmt.Println("变量a的内存地址是: " , &a)
	fmt.Println("指针变量p的存储地址是: " , p)
	fmt.Println("指针变量p的值是: " , *p)
}
