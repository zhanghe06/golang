package main

import (
	"fmt"
	"reflect"
)

// ICommon 某接口
type ICommon interface {
	GetName() string
}

// Animal 第一个结构体，实现了ICommon接口
type Animal struct {
	name string
	age  int64
}

func (u *Animal) GetName() string {
	return u.name
}

func (u *Animal) GetAge() int64 {
	return u.age
}

// Human 第二个结构体，实现了ICommon接口
type Human struct {
	name string
}

func (p *Human) GetName() string {
	return p.name
}

// GetNamesFromList
// 1、先判断一个interface变量是不是数组或切片
// 2、再使用反射和断言确定数组内的元素是不是实现了某个接口
// 3、调用数组成员的接口方法
func GetNamesFromList(list interface{}) []string {
	names := make([]string, 0)

	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(list)
		for i := 0; i < s.Len(); i++ {
			o, ok := s.Index(i).Interface().(ICommon)
			if ok {
				names = append(names, o.GetName())
			}
		}
	}

	return names
}

func main() {
	// 构造两个不同类型的数组
	animals := make([]*Animal, 0)
	animals = append(animals, &Animal{name: "dog"})
	animals = append(animals, &Animal{name: "cat"})
	animals = append(animals, &Animal{name: "bird"})

	humans := make([]*Human, 0)
	humans = append(humans, &Human{name: "XiaoMin"})
	humans = append(humans, &Human{name: "XiaoHong"})
	humans = append(humans, &Human{name: "XiaoQiang"})

	// 使用同一个函数遍历上述两个不同类型的数组
	fmt.Println(GetNamesFromList(animals)) // [dog cat bird]
	fmt.Println(GetNamesFromList(humans))  // [XiaoMin XiaoHong XiaoQiang]
}
