package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//d := time.Now().Add(4 * time.Second) // overslept
	d := time.Now().Add(2 * time.Second) // context deadline exceeded

	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
