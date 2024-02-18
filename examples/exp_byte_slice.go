package main

import (
	"fmt"
)

// 按照指定长度切片分组
func splitSlice(data []byte, sliceLen int) (res [][]byte) {
	sliceIndex := 0
	for {
		minIndex := sliceIndex * sliceLen
		maxIndex := (sliceIndex + 1) * sliceLen
		if maxIndex > len(data) {
			maxIndex = len(data)
		}
		if minIndex >= len(data) {
			break
		}
		item := data[minIndex: maxIndex]
		res = append(res, item)
		sliceIndex++
	}

	return
}

func main() {
	str := "1234567890abcdefghijkl"
	//str := "1234567890abcdefghij"
	//str := ""
	dataByte := []byte(str)
	sliceLen := 5
	res := splitSlice(dataByte, sliceLen)
	fmt.Println(res)
}
