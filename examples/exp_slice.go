package main

import "fmt"

func main() {
	var z []int
	a := make([]int, 5)    // len(a)=5
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	c := b[3:4]
	fmt.Println(z, len(z), cap(z), z == nil) // [] 0 0 true
	fmt.Println(a, len(a), cap(a), a == nil) // [0 0 0 0 0] 5 5 false
	fmt.Println(b, len(b), cap(b), b == nil) // [] 0 5 false
	fmt.Println(c, len(c), cap(c), c == nil) // [0 0 0] 3 5 false

	// 对slice而言,make返回的是对象，所以n := m，就是创建一个对象副本
	m := make([]int, 4)
	m = append(m, 1)
	// m = append(m, 6)

	// 创建一个slice副本, m, n是两个不同的slice实体，内部指针指向同一个底层数组
	n := m

	// 刚开始，mn未扩容时，指向同一个底层数组
	m[1] = 1
	fmt.Println(m) // [0 1 0 0 1]
	fmt.Println(n) // [0 1 0 0 1]

	// 当其中一个变量扩容后，mn开始指向不同的底层数组
	n = append(n, 5)
	fmt.Println(m) // [0 1 0 0 1]
	fmt.Println(n) // [0 1 0 0 1 5]
}
