package main

import (
	"fmt"
	"reflect"
)

type fieldOrder struct {
	ordId      int
	customerId int
}

func createQuery(q interface{}) {
	v := reflect.ValueOf(q)
	if v.Kind() == reflect.Struct {
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}
}

func main() {
	o := fieldOrder{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)
}
