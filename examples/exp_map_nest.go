package main

import (
	"fmt"
)

func testMap() {
	fmt.Println("\ntestMap:")
	var mainMap map[string]string

	mainMap = map[string]string{"1": "a", "2": "b", "3": "c"}

	for key, value := range mainMap {
		fmt.Println(key, ":", value)
	}
}

func testNestMap() {
	fmt.Println("\ntestNestMap:")
	var mainMap = make(map[string]interface{})
	mainMap["1"] = "a"
	mainMap["2"] = "b"
	mainMap["3"] = map[string]string{"1": "a", "2": "b", "3": "c"}

	for key, value := range mainMap {
		fmt.Println(key, ":", value)

		// 断言嵌套结构
		if subMap, ok := value.(map[string]string); ok {
			for k, v := range subMap {
				fmt.Println("\t", k, ":", v)
			}
		}
	}
}

func testMapGet()  {
	params := map[string]interface{}{
		"IF_TYPE":    "1",
		"IT_CONNECTION": []map[string]interface{}{
			{
				"SAP_OBJECT": "BUS2012",
				"OBJECT_ID":  "4500000000",
				"DOCID":      "000C292555BF1EDDAFE545843B7A473F",
				"FILE_IDX":      1,
			},
		},
	}
	res := params["IT_CONNECTION"].([]map[string]interface{})[0]["DOCID"]
	fmt.Println(res)
}

func main() {
	//testMap()
	//testNestMap()
	testMapGet()
}
