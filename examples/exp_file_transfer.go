package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// 读文件
func ftReadFile(filePath string) []byte {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Read error")
	}
	return content
}

// 写文件
func ftWriteFile(filePath string, data []byte) {
	err := ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("Write error")
	}
}

func main() {
	rp := "examples/滴滴电子发票.pdf"
	wp := "examples/滴滴电子发票-w.pdf"

	rb := ftReadFile(rp)
	//fmt.Println(string(c))

	//ftWriteFile(wp, c)

	// byte组合
	header := "x-compId: 123"

	var hb []byte
	wb := bytes.Join([][]byte{
		hb,
		[]byte(header),
		rb,
	}, []byte("\n"))

	ftWriteFile(wp, wb)

	// 查看修改之后的文件内容
	fmt.Println(string(ftReadFile(wp)))
}
