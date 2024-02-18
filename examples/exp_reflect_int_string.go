package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func query(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		sql := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		//args := make([]interface{}, 0)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				//args = append(args, v.Field(i).Int())
				if i == 0 {
					sql = fmt.Sprintf("%s%d", sql, v.Field(i).Int())
				} else {
					sql = fmt.Sprintf("%s, %d", sql, v.Field(i).Int())
				}
			case reflect.String:
				//args = append(args, fmt.Sprintf("\"%s\"", v.Field(i).String()))
				if i == 0 {
					sql = fmt.Sprintf("%s\"%s\"", sql, v.Field(i).String())
				} else {
					sql = fmt.Sprintf("%s, \"%s\"", sql, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		//fmt.Println(args)
		sql = fmt.Sprintf("%s)", sql)
		fmt.Println(sql)
		return

	}
	fmt.Println("unsupported type")
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	query(o) // insert into order values(456, 56)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	query(e) // insert into employee values("Naveen", 565, "Coimbatore", 90000, "India")

	i := 90  // 这是一个错误的尝试
	query(i) // unsupported type
}
