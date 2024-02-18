package main

import "fmt"

func main() {
	var p *int
	fmt.Println("p = ", p)

	var x = 100
	var y = 100
	var xp *int
	xp = &x // 对x做&运算符来获取其地址，然后将该地址分配给指针xp
	fmt.Println("xp = ", xp)
	fmt.Println("*xp = ", *xp) // 要获得指针指向地址的值，我们可以使用*运算符。这叫解引用

	fmt.Println(&x == &x)
	fmt.Println(&x == &y)
}
