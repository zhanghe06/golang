package main

import (
	"fmt"
	"reflect"
)

type OriginStruct struct {
	Id    int `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Age    int `json:"age,omitempty"`
}

func main()  {
	//fields := []string{"Id", "Name"}
	originStruct := OriginStruct{}
	field, ok := reflect.TypeOf(&originStruct).Elem().FieldByName("Id")
	if !ok {
		fmt.Println("reflect failure")
		return
	}
	fmt.Println(field.Name)
	fmt.Println(field.Tag.Get("json"))
}
