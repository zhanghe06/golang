package main

import "fmt"

// LittleAnimal 接口
type LittleAnimal interface {
	Speak() string
}

func PrintAll(variables []interface{}) {
	for _, val := range variables {
		fmt.Println(val)
	}
}

// DoSomething interface{} 作为参数类型，v可以为任何值
func DoSomething(v interface{}) {
	// ...
}

func main() {
	names := []string{"stanley", "david", "oscar"}
	//PrintAll(names)
	variables := make([]interface{}, len(names))
	for i, v := range names {
		variables[i] = v
	}
	PrintAll(variables)
}
