package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"0123456789"

/*
数字随机数
*/
func createRandInt(n int) string {
	format := fmt.Sprintf("%%0%dv", n)
	maxNum := int32(math.Pow(10, float64(n)))
	//fmt.Println(format)
	//fmt.Println(maxNum)
	return fmt.Sprintf(format, rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(maxNum))
}

/*
字符串随机数

6位组合就能达到百亿级别,碰撞概率相对较小
In [1]: 62**6
Out[1]: 56800235584
*/
func createRandString(n int) string {
	return stringWithCharset(n, charset)
}

func stringWithCharset(length int, charset string) string {
	seededRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	fmt.Println(createRandInt(6))
	fmt.Println(createRandString(6))
}
