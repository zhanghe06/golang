package main

import "fmt"

const (
	m = iota // 0
	n = iota // 1
)

const (
	a = "A"  // A
	b = "B"  // B
	c = iota // 2
	d = iota // 3
	e = "E"  // E
	p = iota // 5
	q = iota // 6
	r = iota + 100 // 107
	y // 108
	z // 109
)

// 分组自定义错误码，每组使用const分别定义
const (
	// ErrorCode 通用异常
	ErrorCode = 100000 + iota
	ErrorCodeInternalServerError
)

const (
	// ErrorCodeUser 用户模块
	ErrorCodeUser = 200000 + iota
	ErrorCodeUserNotFound
	ErrorCodeUserDisabled
)

func main() {
	fmt.Println(m, n)
	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(e)
	fmt.Println(p, q)
	fmt.Println(r)
	fmt.Println(y, z)

	fmt.Println(ErrorCode)
	fmt.Println(ErrorCodeInternalServerError)
	fmt.Println(ErrorCodeUser)
	fmt.Println(ErrorCodeUserNotFound)
	fmt.Println(ErrorCodeUserDisabled)
}
