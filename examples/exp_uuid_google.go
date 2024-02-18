package main

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func main() {
	u1, _ := uuid.NewUUID()

	// Creating UUID Version 4
	// panic on error
	//u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)
	fmt.Println(strings.ReplaceAll(u1.String(), "-", ""))
}
