package main

import "fmt"

func testInterface() {
	// 声明并初始化
	var a = []interface{}{123, "hello world", true}
	fmt.Println(a) // [123 hello world true]
	a = append(a, 456)
	fmt.Println(a) // [123 hello world true 456]

	for _, i := range a {
		fmt.Print(i) // 123hello worldtrue456
	}
	fmt.Println()

	// 初始化
	args := make([]interface{}, 0)
	args = append(args, 1, "2", 3, true)
	for _, i := range args {
		fmt.Print(i) // 123true
	}
}

func main() {
	testInterface()
}
