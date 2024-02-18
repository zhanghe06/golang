package main

import (
	"bytes"
	"crypto/aes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)


func padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func secretKeySpec(skey []byte) (key string) {
	//key转换
	has := md5.Sum(skey)
	sky := fmt.Sprintf("%X", has)
	return strings.ToUpper(sky)
}

func AesEncrypt(origData []byte, skey []byte) (string, error) {
	// aes 加密
	skeySpec := secretKeySpec(skey)
	/**AES 加密**/

	//key只能是 16 24 32长度
	block, err := aes.NewCipher([]byte(skeySpec))
	if err != nil {
		return "", err
	}
	//padding
	origData = padding(origData, block.BlockSize())
	//存储每次加密的数据

	//分组分块加密
	buffer := bytes.NewBufferString("")
	tmpData := make([]byte, block.BlockSize()) //存储每次加密的数据
	for index := 0; index < len(origData); index += block.BlockSize() {
		block.Encrypt(tmpData, origData[index:index+block.BlockSize()])
		buffer.Write(tmpData)
	}

	return strings.ToUpper(hex.EncodeToString(buffer.Bytes())), nil
}

func AesDecrypt(text string, skey []byte) (string, error) {
	// aes解密

	skeySpec := secretKeySpec(skey)
	block, err := aes.NewCipher([]byte(skeySpec))
	if err != nil {
		return "", err
	}

	src, _ := hex.DecodeString(text)
	/**AES 解密**/
	buffer := bytes.NewBufferString("")
	tmpData := make([]byte, block.BlockSize())
	for index := 0; index < len(src); index += block.BlockSize() {
		block.Decrypt(tmpData, src[index:index+block.BlockSize()])
		buffer.Write(tmpData)
	}

	// 去掉末尾非打印控制字符
	var deByte []byte
	for i := len(buffer.Bytes()); i > 0; i-- {
		if buffer.Bytes()[i-1] >= 32 {
			deByte = buffer.Bytes()[:i]
			break
		}
	}
	return strings.TrimSpace(string(deByte)), nil
}

func main()  {
	res, err := AesDecrypt("30vq9GpZ/uzK6i3zENjnSw==", []byte("31349862bfa473deba8ephone"))
	fmt.Println(err)
	fmt.Println(res)
}