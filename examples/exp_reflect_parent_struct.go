package main

import (
	"fmt"
	"reflect"
	"strings"
)

// GetReflectField 递归反射结构
func GetReflectField(reflectType reflect.Type, res *[]string) {
	if reflectType.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < reflectType.NumField(); i++ {
		name := reflectType.Field(i).Name
		if reflectType.Field(i).Type.Kind() == reflect.Struct {
			GetReflectField(reflectType.Field(i).Type, res)
			continue
		}
		*res = append(*res, name)
	}
}

// GetReflectTag 递归反射结构
func GetReflectTag(reflectType reflect.Type, res *[]string) {
	if reflectType.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < reflectType.NumField(); i++ {
		tag := reflectType.Field(i).Tag.Get("json")
		if tag == "" {
			GetReflectTag(reflectType.Field(i).Type, res)
			continue
		}
		*res = append(*res, tag)
	}
}

func GetColSQL(model interface{}) {
	var resTag, resName []string

	t := reflect.TypeOf(model)
	// tag
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("json")
		if tag == "" {
			GetReflectTag(t.Field(i).Type, &resTag)
			continue
		}
		resTag = append(resTag, tag)
	}
	// name
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		//if name == "UserA" {
		//	fmt.Println("------")
		//	fmt.Println(t.Field(i).Type.Kind() == reflect.Struct)
		//}
		if t.Field(i).Type.Kind() == reflect.Struct {
			GetReflectField(t.Field(i).Type, &resName)
			continue
		}
		resName = append(resName, name)
	}

	fmt.Println(strings.Join(resTag, ","))
	fmt.Println(strings.Join(resName, ","))
	return
}

func main() {
	type UserA struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	type UserB struct {
		UserA
		Other string `json:"other"`
	}

	type UserC struct {
		a []string
		UserB
		OtherX string `json:"other_x"`
	}

	type UserMore struct {
		UserC
		ShopName string `json:"shop_name"`
	}
	GetColSQL(UserC{})
	//t.Log(GetMoreTableColumnSQL(UserMore{}, []string{"user","shop"}[:]...))
}
