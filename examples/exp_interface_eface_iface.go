package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	// iface 带有方法的接口
	// B
	if live() == nil {
		fmt.Println("A")
	} else {
		fmt.Println("B")
	}

	// eface 空接口
	// C
	var i interface{}
	if i == nil {
		fmt.Println("C")
	}
}
