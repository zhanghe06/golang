package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Req struct {
	Num     int    `json:"num"`
	Age     int64  `json:"age"`
	Year    string `json:"year"`
	Content []byte `json:"content"`
}

func jsonUnmarshal() {
	req := Req{
		Num:     1,
		Age:     12,
		Year:    "2022",
		Content: []byte("123456"),
	}
	var params map[string]interface{}
	reqByte, _ := json.Marshal(req)

	// json.Marshal、json.Unmarshal
	// 数字类型被转为float64
	// 字节类型被Base64编码
	err := json.Unmarshal(reqByte, &params)
	if err != nil {
		return
	}

	fmt.Println(params)
	fmt.Println(params["num"].(float64)) // 需要强转为int
	fmt.Println(params["age"].(float64)) // 需要强转为int64
	fmt.Println(params["content"].(string))

	// base64解码，反转为字节
	content, _ := base64.StdEncoding.DecodeString(params["content"].(string))
	fmt.Println([]byte("123456"))
	fmt.Println(content)
}

//func customUnmarshal()  {
//	var params map[string]interface{}
//	reqByte, _ := json.Marshal(req)
//
//	decoder := json.NewDecoder(bytes.NewReader(reqByte))
//	decoder.UseNumber()
//	err = decoder.Decode(&params)
//	if err != nil {
//		return
//	}
//}

func main() {
	jsonUnmarshal()
	//customUnmarshal()
}
