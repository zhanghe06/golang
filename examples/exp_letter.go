package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Case2Camel 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// UcFirst 首字母大写
func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// LcFirst 首字母小写
func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func main()  {
	fmt.Println(Case2Camel("hello_world")) // HelloWorld
	fmt.Println(UcFirst("hello world")) // Hello world
	fmt.Println(LcFirst("Hello World")) // hello World
}