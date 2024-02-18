package main

import (
	"fmt"
	"time"
)

/*
 * 子协程如果需要使用主协程的变量，需要显示传递
 */
func main() {
	names := []string{"Eric", "Lily", "Tom", "Jim", "Mike"}
	// Eric    Lily    Tom     Jim     Mike
	for _, name := range names {
		func() {
			fmt.Printf("%s\t", name)
		}()
	}

	time.Sleep(time.Second)
	fmt.Printf("\n")

	// Mike    Mike    Mike    Mike    Mike
	for _, name := range names {
		go func() {
			// Loop variables captured by 'func' literals in 'go' statements might have unexpected values
			fmt.Printf("%s\t", name)
		}()
	}

	time.Sleep(time.Second)
	fmt.Printf("\n")

	// Eric    Mike    Tom     Jim     Lily
	for _, name := range names {
		go func(name string) {
			fmt.Printf("%s\t", name)
		}(name)
	}
	time.Sleep(time.Second)
}
