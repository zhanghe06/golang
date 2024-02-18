package main

import (
	"fmt"
	"strings"
)

// 多态
var t interface {
	talk() string
}

type martian struct {
}

func (m martian) talk() string {
	return "ok"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

// 形参
type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

// 嵌入
type starship struct {
	laser
}

func main() {
	// 以下演示接口的3种常用用法

	// 多态
	t = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())

	// 形参
	shout(martian{})
	shout(laser(2))

	// 嵌入
	s := starship{laser(3)}
	fmt.Println(s.talk())
	shout(s)
}
