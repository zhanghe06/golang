package main

import "fmt"

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		case []interface{}:
			fmt.Printf("Param #%d is a array\n", i)
		case map[string]interface{}:
			fmt.Printf("Param #%d is a map\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}

func main() {
	classifier(1, "1", true, map[string]string{"1": "a"})
}
