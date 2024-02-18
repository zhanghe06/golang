package main

import (
	"fmt"
	"reflect"
)

// 定义结构体
type User struct {
	Id   int
	Name string
	Age  int
}

// 匿名字段
type Boy struct {
	User
	Addr string
}

func main() {
	m := Boy{User{1, "zs", 20}, "bj"}
	t := reflect.TypeOf(m)
	fmt.Println(t) // main.Boy
	// Anonymous：匿名
	fmt.Printf("%#v\n", t.Field(0)) // reflect.StructField{Name:"User", PkgPath:"", Type:(*reflect.rtype)(0x109a6c0), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}
	// 值信息
	fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0)) // main.User{Id:1, Name:"zs", Age:20}
}
