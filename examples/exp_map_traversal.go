package main

import "fmt"

const (
	Year = iota
	Last
)

func main() {
	testMap := make(map[string]interface{})
	testMap["name"] = "tom"
	testMap["host"] = "dev"


	// Copy Map
	copyMap := make(map[string]interface{})
	for key, value := range testMap {
		copyMap[key] = value
	}

	testMap["port"] = 8000
	testMap["year"] = 0

	fmt.Println(len(testMap))

	for k, v := range testMap {
		fmt.Printf("%s: %v\n", k, v)
	}

	y, ok := testMap["year"]
	fmt.Println(ok)
	fmt.Println(y == Year)
	fmt.Println(Last)
	if ok {
		fmt.Println("+++")
	} else {
		fmt.Println("---")
	}

	for k, v := range copyMap {
		fmt.Printf("%s: %v\n", k, v)
	}
}
