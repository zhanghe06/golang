package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayhelloGolang(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("path", r.URL.Path)
	w.Write([]byte("Hello Golang"))
}

func main() {
	http.HandleFunc("/", sayhelloGolang)     //设置访问的路由
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
