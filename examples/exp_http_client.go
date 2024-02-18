package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func clientGet()  {
	res, err := http.Get("http://127.0.0.1:8080")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer res.Body.Close()
	resByte, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(resByte))
}

func clientGetWithHeader()  {
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8080", nil)
	req.Header.Set("token", "00998ecf8427e")
	res, err := (&http.Client{}).Do(req)

	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer res.Body.Close()
	resByte, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(resByte))
}

func clientPost()  {
	type User struct {
		Id int32
		Name string
	}
	userInfo := User {
		1,
		"name",
	}
	userInfoByte, _ := json.Marshal(userInfo)

	res, err := http.Post("http://127.0.0.1:8080", "application/json", strings.NewReader(string(userInfoByte)))
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer res.Body.Close()
	resByte, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(resByte))
}


func main() {
	clientGet()
	clientGetWithHeader()
	clientPost()
}
