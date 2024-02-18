package main

import (
	"errors"
	"fmt"
)

func errorEmpty() (err error) {
	return
}

func errorFill() (err error) {
	err = errors.New("test")
	return
}

func main() {
	err := errorEmpty()
	fmt.Printf("errorEmpty: %v\n", err)

	err = errorFill()
	fmt.Printf("errorFill: %v\n", err)
}
