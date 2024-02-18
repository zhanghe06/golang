package main

import "fmt"

// 集合（特性：交集、差集、并集）

// 取两个切片的交集
func intersect(slice1, slice2 []interface{}) []interface{} {
	m := make(map[interface{}]int)
	n := make([]interface{}, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times := m[v]
		if times >= 1 {
			n = append(n, v)
		}
	}
	return n
}

// 取两个切片的差集
func difference(slice1, slice2 []interface{}) []interface{} {
	m := make(map[interface{}]int)
	n := make([]interface{}, 0)
	inter := intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}
	for _, value := range slice1 {
		if m[value] == 0 {
			n = append(n, value)
		}
	}
	return n
}

// 取两个切片的并集
func union(slice1, slice2 []interface{}) []interface{} {
	m := make(map[interface{}]int)
	n := make([]interface{}, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		m[v]++
	}
	for k := range m {
		n = append(n, k)
	}
	return n
}

func main() {
	sa := make([]interface{}, 0)
	sb := make([]interface{}, 0)
	for _, v := range []int{2, 3, 4, 3} {
		sa = append(sa, v)
	}
	for _, v := range []int{1, 2, 3} {
		sb = append(sb, v)
	}
	// 交集
	fmt.Println(intersect(sa, sb))
	// 差集
	fmt.Println(difference(sa, sb))
	fmt.Println(difference(sb, sa))
	// 并集
	fmt.Println(union(sa, sb))
}
