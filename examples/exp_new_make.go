package main

import "fmt"

func main() {
	// 初始化 整数
	i1 := new(int)
	*i1 = 1
	fmt.Println(i1) // 0xc0000160b8

	i2 := new(*int)
	i2 = &i1
	fmt.Println(i2)  // 0xc00000e028
	fmt.Println(*i2) // 0xc0000160b8

	// 初始化 数组
	arrayFirst := new([5]int)
	fmt.Println(arrayFirst) // &[0 0 0 0 0]

	// 初始化并赋值 数组
	arraySecond := [5]int{1, 2, 3} // 定长
	fmt.Println(arraySecond)       // [1 2 3 0 0]

	arrayThird := [...]int{6, 7, 8} // 不定长
	fmt.Println(arrayThird)         // [6 7 8]

	arrayFourth := [5]string{3: "Chris", 4: "Ron"} // 初始化其中的部分元素
	fmt.Println(arrayFourth)                       // [   Chris Ron]

	// 初始化 slice（容器类型）
	sliceFirst := make([]int, 5)
	fmt.Println(sliceFirst) // [0 0 0 0 0]

	sliceSecond := new([]int)
	fmt.Println(sliceSecond) // &[]

	// 初始化 map（容器类型）
	m := make(map[string]int)
	fmt.Println(m) // map[]

	n := new(map[string]int)
	fmt.Println(n) // &map[]

	// 初始化 struct
	type Show struct{}
	s1 := new(Show)
	fmt.Println(s1) // &{}

	s2 := &Show{}
	fmt.Println(s2) // &{}

	// 初始化 channel
	c1 := make(chan bool) // 无缓冲
	fmt.Println(c1)       // 0xc000060060

	c2 := make(chan int, 1) // 有缓冲
	fmt.Println(c2)         // 0xc000050070
}
