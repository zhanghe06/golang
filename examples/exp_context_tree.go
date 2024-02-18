package main

import (
	"context"
	"fmt"
	"time"
)

func tree() {
	ctx1 := context.Background()
	fmt.Println("ctx1: ", ctx1)
	ctx2, cancel := context.WithCancel(ctx1)
	defer cancel()
	fmt.Println("ctx2: ", ctx2)
	ctx3, cancel := context.WithTimeout(ctx2, time.Second * 5)
	defer cancel()
	fmt.Println("ctx3: ", ctx3)
	ctx4, cancel := context.WithTimeout(ctx3, time.Second * 3)
	defer cancel()
	fmt.Println("ctx4: ", ctx4)
	ctx5, cancel := context.WithTimeout(ctx3, time.Second * 6)
	defer cancel()
	fmt.Println("ctx5: ", ctx5)
	ctx6 := context.WithValue(ctx5, "userID", 12)
	fmt.Println("ctx6: ", ctx6)
}


func main() {
	tree()
	time.Sleep(time.Second*3)
}
