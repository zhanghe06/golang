package main

import "fmt"
import "encoding/binary"

func main() {
	a := []byte{0, 1, 2, 3}
	fmt.Println(a)  // [0 1 2 3]
	fmt.Println(binary.BigEndian.Uint32(a))  // 66051
	fmt.Println(binary.LittleEndian.Uint32(a))  // 50462976
}

// byte数组如何转int，两种方式：大端和小端
// 大端就是内存中低地址对应着整数的高位。
// 小端就是反过来。
// 1个字节8位
// 大端：00000000 00000001 00000010 00000011
// 小端：00000011 00000010 00000001 00000000

// Golang encoding/binary.go 没有对byte数组长度进行检查
// 如果传入的数组长度小于4，自然会报错：panic: runtime error: index out of range。
