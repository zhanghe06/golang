package main

import "strconv"

func main() {
	i, err := strconv.Atoi("12345")
	if err != nil {
		panic(err)
	}
	i *= 10
	println(i)
}
