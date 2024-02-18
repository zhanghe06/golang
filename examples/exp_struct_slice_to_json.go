package main

import "fmt"
import "encoding/json"

func main() {
	type User struct {
		Id    int    `json:"id"`
		Name string `json:"name"`
	}

	var userList []*User

	userList = append(userList, &User{1, "A"})
	userList = append(userList, &User{2, "B"})

	userListJson, err := json.Marshal(userList) //转换成JSON返回的是byte[]
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(userListJson))

}
