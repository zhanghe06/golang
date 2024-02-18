package main

import "fmt"
import "encoding/json"

func main() {
	type TestStructNest struct {
		DataString string `json:"dataString"`
		DataInt    int    `json:"dataInt"`
	}

	type TestStruct struct {
		DataStruct          struct{}         `json:"dataStruct"`
		DataString          string           `json:"dataString"`
		DataInt             int              `json:"dataInt"`
		DataStringPoint     *string          `json:"dataStringPoint"`
		DataIntPoint        *int             `json:"dataIntPoint"`
		DataStringArray     []string         `json:"dataStringArray"`
		DataStructArray     []struct{}       `json:"dataStructArray"`
		DataStructNest      TestStructNest   `json:"dataStructNest"`
		DataStructNestArray []TestStructNest `json:"dataStructNestArray"`
		DataStringOmitempty string           `json:"dataStringOmitempty,omitempty"`
	}

	// 空值结构(忽略omitempty序列化字段)
	testStructEmpty := &TestStruct{
		DataStringArray: make([]string, 0),
	}
	jsonBytesEmpty, err := json.Marshal(testStructEmpty) //转换成JSON返回的是byte[]
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(jsonBytesEmpty))

	// 零值结构(忽略omitempty序列化字段)
	jsonBytesZero, _ := json.Marshal(&TestStruct{
		DataStringOmitempty: "",
	}) //转换成JSON返回的是byte[]
	fmt.Println(string(jsonBytesZero))

	// 非空结构
	jsonBytesNew, _ := json.Marshal(&TestStruct{
		DataStringOmitempty: "123",
	}) //转换成JSON返回的是byte[]
	fmt.Println(string(jsonBytesNew))
}
