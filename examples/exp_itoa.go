package main

import "strconv"

func main() {
	s := strconv.Itoa(12345)
	s += "0"
	println(s)
}
