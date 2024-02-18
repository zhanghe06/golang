package main

import (
	"fmt"
	"sort"
)

type StructSortPerson struct {
	Name string
	Age  int
}

type ByAge []StructSortPerson

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func reverseSlice(s []StructSortPerson) {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
}

func main()  {
	people := []StructSortPerson{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	sort.Sort(ByAge(people))

	fmt.Println(people)

	reverseSlice(people)

	fmt.Println(people)
}


