package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// 匿名嵌套结构体，初始化时，需要指定结构体类型
	type T struct {
		User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		} `json:"user"`
	}
	req := T{
		struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{
			"Tom",
			12,
		},
	}

	var params map[string]interface{}
	reqByte, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(reqByte, &params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(params)
}
