package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS7UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func ECBEncrypt(block cipher.Block, src, key []byte) ([]byte, error) {
	blockSize := block.BlockSize()

	encryptData := make([]byte, len(src))
	tmpData := make([]byte, blockSize)

	for index := 0; index < len(src); index += blockSize {
		block.Encrypt(tmpData, src[index:index+blockSize])
		copy(encryptData, tmpData)
	}
	return encryptData, nil
}

func ECBDecrypt(block cipher.Block, src, key []byte) ([]byte, error) {
	dst := make([]byte, len(src))

	blockSize := block.BlockSize()
	tmpData := make([]byte, blockSize)

	for index := 0; index < len(src); index += blockSize {
		block.Decrypt(tmpData, src[index:index+blockSize])
		copy(dst, tmpData)
	}

	return dst, nil
}

func main() {

	// 加密
	src := []byte("15921082801")
	key := []byte("31349862bfa473deba8ephon")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	src = PKCS7Padding(src, block.BlockSize())

	dst, err := ECBEncrypt(block, src, key)
	if err != nil {
		panic(err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(dst)) // SpfAShHImQhWjd/21Pgz2Q==


	//data, err := base64.StdEncoding.DecodeString("30vq9GpZ/uzK6i3zENjnSw==")

	// 解密
	//src, err = ECBDecrypt(block, dst, key)
	src, err = ECBDecrypt(block, dst, key)
	if err != nil {
		panic(err)
	}

	src = PKCS7UnPadding(src)

	fmt.Println(string(src)) // 123456

}