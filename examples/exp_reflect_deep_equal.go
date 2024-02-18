package main

import (
	"fmt"
	"reflect"
)

type reflectS struct {
	Name    string
	Age     int
	Address *int
	Data    []int
}

func main() {
	a := reflectS{
		Name:    "aa",
		Age:     1,
		Address: new(int),
		Data:    []int{1, 2, 3},
	}
	b := reflectS{
		Name:    "aa",
		Age:     1,
		Address: new(int),
		Data:    []int{1, 2, 3},
	}

	fmt.Println(reflect.TypeOf(a))                           // main.reflectS
	fmt.Println(reflect.TypeOf(a).Kind())                    // struct
	fmt.Println(reflect.TypeOf(a).Kind() == reflect.Struct)  // true
	fmt.Println(reflect.ValueOf(a))                          // {aa 1 0xc0000180b8 [1 2 3]}
	fmt.Println(reflect.ValueOf(a).Kind())                   // struct
	fmt.Println(reflect.ValueOf(a).Kind() == reflect.Struct) // true
	//fmt.Println(reflect.Kind(a))
	// fmt.Println(a == b)  // 实例不能比较 Invalid operation: a == b (the operator == is not defined on reflectS)
	fmt.Println(&a == &b)                // 指针可以比较 false
	fmt.Println(reflect.DeepEqual(a, b)) // true
}
