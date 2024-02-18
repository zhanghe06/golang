package main

import "fmt"

// ...
// 1022 1024
// 1024 1024
// 2 1024
// 4 1024
// ...
func main() {
	sumBytes := make([]byte, 0, 1024)
	subBytes := make([]byte, 2, 2)
	for {
		select {
		default:
			// fmt.Println(subBytes)
			sumBytes = append(sumBytes, subBytes...)
			fmt.Println(len(sumBytes), cap(sumBytes))
			if len(sumBytes) == cap(sumBytes) {
				sumBytes = sumBytes[:0]
			}
		}
	}
}
