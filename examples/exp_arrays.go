package main

import "fmt"

func testArray01() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func testArray02() {
	var a [4]int           // 元素自动初始化为零[0 0 0 0]
	b := [4]int{2, 5}      // 未提供初始化值得元素自动初始化为0  [2 5 0 0]
	c := [...]int{1, 2, 3} // 编译器按初始化值数量确定数组长度 [1 2 3]
	d := [2]string{"Hello", "World"}
	e := [...]int{10, 3: 100}      // 支持索引初始化，但注意数组长度与此有关 [10 0 0 100]
	f := [4]string{1: "b", 3: "c"} // 可以指定初始化的位置

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

func testArray03() {
	// 数组添加元素（数组本身不能添加元素，因为golang中数组长度不可变，通过切片可以）
	a := []int{1, 2, 4, 5, 6}
	i := 2
	v := 3
	b := append([]int{v}, a[i:]...) // [3 4 5 6]
	a = append(a[0:i], b...)
	fmt.Println(a)
}

func testArray04() {
	// 数组删除元素（同理，使用切片操作数组）
	a := []int{0, 1, 2, 3, 4}
	// 删除第i个元素
	i := 2
	a = append(a[:i], a[i+1:]...)
	fmt.Println(a)
}

func main() {
	//testArray01()
	//testArray02()
	testArray03()
	//testArray04()
}
