package main

import "fmt"

func main() {
	str := "09a53916-8a78-4498-bc74-fca2d2b0adbc-xxxxxx"
	//content := str[0:len(str)-1]
	content := str[:36]
	fmt.Println(content)
}
