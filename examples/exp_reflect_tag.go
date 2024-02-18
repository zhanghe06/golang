package main

import (
	"fmt"
	"reflect"
)

type TagStruct struct {
	Name    string `json:"name,omitempty"`
	Age     int
}

func main()  {
	tagStruct := TagStruct{}
	field, ok := reflect.TypeOf(&tagStruct).Elem().FieldByName("Name")
	if !ok {
		fmt.Println("reflect failure")
		return
	}
	fmt.Println(field.Name)
	fmt.Println(field.Tag.Get("json"))
}
