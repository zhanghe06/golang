package main

import "fmt"

func printSliceArgs(v ...interface{}) {
	fmt.Println(v) // 切片，例如: [1 2 3]
	fmt.Println(v...) // 切片打散，例如: 1 2 3
}

func main()  {
	printSliceArgs(1, 2, 3)



	var luckyBoxes [8]int
	var emptyBoxes [36]int
	luckyBoxes = [8]int{1}
	emptyBoxes = [36]int{1, 1, 1, 1, 1, 1, 1, 1}

	fmt.Println(luckyBoxes) // [1 0 0 0 0 0 0 0]
	fmt.Println(emptyBoxes) // [1 1 1 1 1 1 1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
}
