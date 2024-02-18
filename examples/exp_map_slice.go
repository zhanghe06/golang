package main

import "fmt"

type student struct {
	Name string
	Age  int
}

var students = []student{
	{Name: "zhou", Age: 24},
	{Name: "li", Age: 23},
	{Name: "gao", Age: 22},
}

func parseStudentFailure() {
	m := make(map[string]*student)
	for _, stu := range students {
		m[stu.Name] = &stu // 将临时变量stu的地址给map，最后所有的地址都为同样的一个值
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func parseStudentSuccess() {
	m := make(map[string]*student)
	for i, stu := range students {
		m[stu.Name] = &students[i]
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main() {
	// li &{gao 22}
	// gao &{gao 22}
	// zhou &{gao 22}
	parseStudentFailure()

	// zhou &{zhou 24}
	// li &{li 23}
	// gao &{gao 22}
	parseStudentSuccess()

	// 现象分析
	// 遍历获取的是值的拷贝，要修改数组的值还是要用下标

	// 另外，golang里面的map，当通过key获取到value时，这个value是不可寻址的。
	// 因为map会进行动态扩容，当进行扩展后，map的value就会进行内存迁移，其地址发生变化，所以无法对这个value进行寻址。
}
