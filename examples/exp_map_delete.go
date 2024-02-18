package main

import "fmt"

func main()  {
	var m1 map[string]interface{}

	m1 = make(map[string]interface{})
	m1["code"] = 1
	m1["name"] = "1"

	delete(m1, "name")

	var m2 map[string]interface{}

	m2 = make(map[string]interface{})
	m2["code"] = 2
	m2["name"] = "2"
	m2["last"] = []int{1, 2, 3}
	m2["next"] = []map[string]interface{}{m1}

	delete(m2, "last")
	delete(m2, "next")

	fmt.Println(m2)
}