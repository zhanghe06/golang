package main

import (
	"fmt"
	"reflect"
)

func main() {

	v1 := "123456"
	v2 := 12

	fmt.Printf("v1 type: %T\n", v1) // v1 type: string
	fmt.Printf("v2 type: %T\n", v2) // v2 type: int

	// reflect
	fmt.Println("v1 type:", reflect.TypeOf(v1)) // v1 type: string
	fmt.Println("v2 type:", reflect.TypeOf(v2)) // v2 type: int

	// 使用类型 switch 判断
	switch i := interface{}(v).(type) {
	case int, int8, int16, int32, int64:
		fmt.Printf("变量 %d 的类型为 %T. \n", i, v)
	case uint, uint8, uint16, uint32, uint64:
		fmt.Printf("变量 %d 的类型为 %T. \n", i, v)
	default:
		fmt.Printf("%v不是匹配的类型", i)
	}
}
