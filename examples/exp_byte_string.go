package main

import (
	"fmt"
)

func main() {
	str := "hello"
	dataByte := []byte(str)
	fmt.Println(dataByte)
	dataString := string(dataByte[:])
	fmt.Println(dataString)
}
