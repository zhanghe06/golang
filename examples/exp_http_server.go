package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	defer fmt.Fprintf(w, "ok\n")

	fmt.Println("method:", r.Method)
	fmt.Println("method:", r.URL)
	fmt.Println("header:", r.Header)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	println("json:", string(body))

	type User struct {
		Id int32
		Name string
	}
	var user User
	if err = json.Unmarshal(body, &user); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return
	}
	fmt.Printf("%+v", user)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
