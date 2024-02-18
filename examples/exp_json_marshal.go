package main

import (
	"encoding/json"
	"fmt"
)

type X struct {
	B json.RawMessage `json:"b"`
}

func main() {
	x := X{B: []byte(`{"test":"t"}`)}

	b, err := json.Marshal(x)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%s\n", b)

	b, err = json.Marshal(&x)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%s\n", b)

	return
}
