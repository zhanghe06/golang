package main

import "fmt"

// SliceRemoveItem 移位法（修改原切片）
func SliceRemoveItem(a []string, elem string) []string {
	j := 0
	for _, v := range a {
		if v != elem {
			a[j] = v
			j++
		}
	}
	return a[:j]
}

func main() {
	s := []string{"a", "b", "c", "d"}
	fmt.Println(SliceRemoveItem(s, "c"))
	fmt.Println(s)
}
