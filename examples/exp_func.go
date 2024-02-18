package main

import "fmt"

func func1()  {
	fmt.Println("func1")
}

func func2()  {
	fmt.Println("func2")
	return
}

func func3()  {
	fmt.Println("func2")
}

func run()  {
	func1()
	func2()
	func3()
}

func main()  {
	run()
}
