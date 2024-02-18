package main

import (
	"fmt"
	"time"
)

var token string

func refreshToken()  {
	for {
		token = fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(5 * time.Second)
	}
}

func getToken()  {
	fmt.Println(token)
}

func main() {
	go refreshToken()
	for {
		getToken()
		time.Sleep(1 * time.Second)
	}
}
