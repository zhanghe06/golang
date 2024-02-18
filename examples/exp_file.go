package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(filePath string) []byte {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Read error")
	}
	return content
}

func main() {
	// 文件读取
	c := readFile("/etc/profile")
	fmt.Println(string(c))
}
