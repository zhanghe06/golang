package main

import "fmt"

// Container 接口
type Container interface {
	Fuck(msg string) (err error)
}

// container 结构体
type container struct{}

// Fuck container 结构体实现接口
func (c *container) Fuck(msg string) (err error) {
	fmt.Println(msg)
	return err
}

func main() {
	// 判断接口是否已经实现 方式一
	var _ = (Container)(&container{})
	// 判断接口是否已经实现 方式二
	var _ Container = (*container)(nil)
	// 判断接口是否已经实现 方式三
	var _ Container = &container{}
}
