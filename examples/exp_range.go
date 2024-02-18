package main

import "fmt"

func main() {
	// range map

	// 1. 声明 + 创建 + 赋值
	var m map[string]int
	m = make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	// 2. 声明、创建一体 + 赋值
	// m := make(map[string]int)
	// m["a"] = 1
	// m["b"] = 2
	// m["c"] = 3

	// 声明、创建、赋值一体化
	// m := map[string]int{
	//    "a": 1,
	//    "b": 2,
	//    "c": 3,
	// }

	for k, v := range m {
		fmt.Printf("%s: %d\n", k, v)
	}

	// range slice
	var s = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range s {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	//fmt.Println([]int{}[:0])
}
