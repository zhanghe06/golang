package main

import (
	"fmt"
	"github.com/agclqq/goencryption"
)

// 在线验证：https://tool.lmeee.com/jiami/des3

// https://blog.csdn.net/agclqq/article/details/119878572
// 3DES加密的密钥（key）必须要24个字节

// https://tech1024.com/original/3015
// https://github.com/forgoer/openssl

// 3DES ECB pkcs7padding base64
func ecbDecrypt() {
	easyType := "3des/ecb/pkcs7/base64"
	cipherText := "30vq9GpZ/uzK6i3zENjnSw=="
	key := "31349862bfa473deba8ephone"
	iv := ""

	key = key[:24]

	text, err := goencryption.EasyDecrypt(easyType, cipherText, key, iv)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(text)
}

func main() {
	ecbDecrypt()
}
