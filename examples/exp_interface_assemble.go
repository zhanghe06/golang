package main

import "fmt"

type DevUser struct {
	name string
}

type CallBack interface {
	getName() string
	setName(string)
	BaseCall
}

type BaseCall interface {
	doSomething()
}

func (user DevUser) doSomething() {
	fmt.Println("something to do")
}

func (user DevUser) getName() string {
	return user.name
}

func (user *DevUser) setName(name string) {
	user.name = name
}

// 强制转换，判断User结构体是否实现了CallBack接口
func hasImplementCallBackInterface(callback CallBack) {
	if u, ok := callback.(*DevUser); ok {
		fmt.Println("name is:", u.name)
		return
	}
	fmt.Println("not has interface")
}

// 通用判断，判断某结构体是否实现了接口
func hasImplementInterface(instance interface{}) {
	switch t := instance.(type) {
	case DevUser:
		fmt.Println("struct is:", t.name)
	case BaseCall:
		fmt.Println("struct is: BaseCall")
	default:
		fmt.Println("unknow")
	}
}

func main() {
	user := DevUser{}
	user.doSomething()
	user.setName("tom")
	fmt.Println(user.getName())
	hasImplementCallBackInterface(&user)
	hasImplementInterface(user)
}
