package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

const IntSize = int(unsafe.Sizeof(0))

//判断我们系统中的字节序类型
func systemEndian() {
	var i = 0x1
	bs := (*[IntSize]byte)(unsafe.Pointer(&i))
	if bs[0] == 0 {
		fmt.Println("system endian is little endian")
	} else {
		fmt.Println("system endian is big endian")
	}
}

func testBigEndian() {

	// 0000 0000 0000 0000   0000 0001 1111 1111
	var testInt int32 = 256
	fmt.Printf("%d use big endian: \n", testInt)
	var testBytes = make([]byte, 4)
	binary.BigEndian.PutUint32(testBytes, uint32(testInt))
	fmt.Println("int32 to bytes:", testBytes)

	resInt := binary.BigEndian.Uint32(testBytes)
	fmt.Printf("bytes to int32: %d\n\n", resInt)
}

func testLittleEndian() {

	// 0000 0000 0000 0000   0000 0001 1111 1111
	var testInt int32 = 256
	fmt.Printf("%d use little endian: \n", testInt)
	var testBytes = make([]byte, 4)
	binary.LittleEndian.PutUint32(testBytes, uint32(testInt))
	fmt.Println("int32 to bytes:", testBytes)

	resInt := binary.LittleEndian.Uint32(testBytes)
	fmt.Printf("bytes to int32: %d\n\n", resInt)
}

func main() {
	systemEndian()
	fmt.Println("")
	testBigEndian()
	testLittleEndian()
}

//system edian is big endian
//
//256 use big endian:
//int32 to bytes: [0 0 1 0]
//bytes to int32: 256
//
//256 use little endian:
//int32 to bytes: [0 1 0 0]
//bytes to int32: 256
