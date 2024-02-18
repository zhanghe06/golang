package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	f, err := os.Open("examples/test.txt")
	if err != nil {
		return
	}
	defer f.Close()

	buffer := make([]byte, 512)
	_, err = f.Read(buffer)
	if err != nil {
		return
	}
	contentType := http.DetectContentType(buffer)
	fmt.Println(contentType)
}
