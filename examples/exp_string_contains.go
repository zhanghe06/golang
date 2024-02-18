package main

import (
	"fmt"
	"strings"
)

// 语句中是否包含所有指定关键字
func testStringContains(sentence, words, separator string) bool {
	keywords := strings.Split(words, separator)
	result := true
	for _, v := range keywords {
		result = result && strings.Contains(sentence, strings.Trim(v, " "))
	}
	fmt.Println(result)
	return result
}

func main() {
	testStringContains("I like golang and python", "golang, python", ",")
	testStringContains("I like golang and python", "python, golang", ",")
	testStringContains("I like golang and python", "python, golang, js", ",")
	testStringContains("I like golang and python", "golang, js", ",")
	testStringContains("I like golang and python", "js", ",")
	testStringContains("I like golang and python", "", ",")
}
