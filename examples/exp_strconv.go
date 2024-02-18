package main

import (
	"fmt"
	"strconv"
)

func main()  {
	fmt.Printf("%T\n", strconv.Itoa(1))

	i, _ := strconv.Atoi("1")
	fmt.Printf("%T\n", i)

	fmt.Printf("%T\n", strconv.FormatInt(int64(1), 10))

	fmt.Printf("%T\n", strconv.FormatUint(uint64(1), 10))
}


//string转成int
//int, err := strconv.Atoi(string)

//string转成int64
//int64, err := strconv.ParseInt(string, 10, 64)

//int转成string
//string := strconv.Itoa(int)

//int64转成string
//string := strconv.FormatInt(int64, 10)
