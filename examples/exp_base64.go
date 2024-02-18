package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	// 编码
	srcE := []byte("EUq1vfI1PTRM3dXO")
	dstE := base64.StdEncoding.EncodeToString(srcE)
	fmt.Println(dstE)

	// 解码
	srcD := "RVVxMXZmSTFQVFJNM2RYTw=="
	dstD, err := base64.StdEncoding.DecodeString(srcD)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dstD))
}
