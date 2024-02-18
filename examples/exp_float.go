package main

import "fmt"

func main()  {
	var a interface{}
	var b int

	a = 1
	a = float64(1)
	b = 1

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a == b)
	fmt.Println(a == float64(b))
}