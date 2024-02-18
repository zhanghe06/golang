package main

import (
	"reflect"
)

func getStructFields(s interface{}) (res []string) {
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		res = append(res, t.Field(i).Name)
	}
	return
}

func main()  {
	type Employee struct {
		Name    string
		Id      int
		Address string
		salary  int
		Country string
	}
	getStructFields(Employee{})
}
