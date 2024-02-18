package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//Write buf is: [1 0]
//Put l buf is: [1 0]
//Get l buf is: 1
//Put b buf is: [0 1]
//Get b buf is: 1
func main() {
	int16bufW := new(bytes.Buffer)
	_ = binary.Write(int16bufW, binary.LittleEndian, uint16(1))
	fmt.Println("Write buf is:", int16bufW.Bytes())

	var int16bufL [2]byte
	// 通过调用binary.LittleEndian.PutUint16,可以按照小端序的格式将uint16类型的数据序列化到buffer中
	binary.LittleEndian.PutUint16(int16bufL[:], uint16(1))
	fmt.Println("Put l buf is:", int16bufL[:])

	// 通过binary.LittleEndian.Uint16将buffer中内容反序列化出来
	iL := binary.LittleEndian.Uint16(int16bufL[:])
	fmt.Println("Get l buf is:", iL)

	var int16bufB [2]byte
	// 通过调用binary.BigEndian.PutUint16,可以按照小端序的格式将uint16类型的数据序列化到buffer中
	binary.BigEndian.PutUint16(int16bufB[:], uint16(1))
	fmt.Println("Put b buf is:", int16bufB[:])

	// 通过binary.BigEndian.Uint16将buffer中内容反序列化出来
	iB := binary.BigEndian.Uint16(int16bufB[:])
	fmt.Println("Get b buf is:", iB)
}
