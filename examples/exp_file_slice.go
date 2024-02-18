package main

import (
	"fmt"
	"os"
	"time"
)

// FIXME TODO 2023-03-14

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func cat(f *os.File) []byte {
	var payload []byte
	for {
		buf := make([]byte, 1024)
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // EOF
			return payload
		case nr > 0:
			payload = append(payload, buf...)
		}
	}
}


func main() {
	file, err := os.Open("test.flv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println(file)
	payload := cat(file)
	fo, errs := os.Create(fmt.Sprintf("./%d.bmp", time.Now().UnixNano())) //time.Now().UnixNano()
	check(errs)
	_, errs = fo.Write(payload)
}
