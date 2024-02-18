package main

import "fmt"

// NaturalPerson 定义结构体
type NaturalPerson struct {
	Country string
}

// YoungStudent 定义结构体并继承 NaturalPerson
type YoungStudent struct {
	Name          string
	Age           int
	Gender        string
	NaturalPerson // 继承 NaturalPerson 所有属性与方法
}

// 为 YoungStudent 结构体添加方法 打印姓名
func (s *YoungStudent) fmtName() {
	fmt.Println(s.Name)
}

// 为 NaturalPerson 结构体添加方法 打印国家
func (p *NaturalPerson) fmtCountry() {
	fmt.Println(p.Country)
}

func main() {
	ys := new(YoungStudent)
	// 这是 NaturalPerson 的属性与方法
	ys.Country = "中国"
	ys.fmtCountry()
	// 这是 YoungStudent 的属性与方法
	ys.Name = "王二狗"
	ys.fmtName()
}
