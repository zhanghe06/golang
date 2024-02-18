package main

import "fmt"

func assert01()  {
	// 接口类型的类型断言
	var i interface{}
	i = "hello"

	s := i.(string)
	fmt.Println(s) // hello

	s, ok := i.(string)
	fmt.Println(s, ok) // hello true

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false

	i = 100
	t, ok := i.(int)
	fmt.Println(t, ok) // 100 true

	t2 := i.(string)
	fmt.Println(t2) // panic
}

func assert02()  {
	// 非接口类型的类型断言
	var i int
	i = 100
	//if j, ok := interface{}(i).(int); ok {
	//	fmt.Println(j)
	//}
	j, ok := interface{}(i).(int64)
	fmt.Println(j, ok) // 0, false
}

func main() {
	//assert01()
	assert02()
}
