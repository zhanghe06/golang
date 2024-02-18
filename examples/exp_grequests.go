package main

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"log"
)

func getSession() {
	session := grequests.NewSession(nil)
	resp, err := session.Get(
		"http://httpbin.org/cookies/set",
		&grequests.RequestOptions{Params: map[string]string{"one": "two"}},
	)
	if err != nil {
		log.Fatal("Cannot set cookie: ", err)
	}
	if resp.Ok != true {
		log.Println("Request did not return OK")
	}
	fmt.Println(resp.String())
}

func get() {
	ro := &grequests.RequestOptions{Params: map[string]string{"One": "Two"}}
	resp, err := grequests.Get("http://httpbin.org/get", ro)

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	// 打印返回内容
	fmt.Println(resp.String())

	type GetResponse struct {
		Args    struct{} `json:"args"`
		Headers struct {
			Accept         string `json:"Accept"`
			AcceptEncoding string `json:"Accept-Encoding"`
			AcceptLanguage string `json:"Accept-Language"`
			Dnt            string `json:"Dst"`
			Host           string `json:"Host"`
			UserAgent      string `json:"User-Agent"`
			Authorization  string `json:"Authorization"`
		} `json:"headers"`
		Origin string `json:"origin"`
		URL    string `json:"url"`
	}
	getResponse := &GetResponse{}
	if err := resp.JSON(getResponse); err != nil {
		fmt.Println(err.Error())
	}
	// 打印构造结构
	fmt.Println(getResponse)
	fmt.Println(getResponse.URL)

	// 打印转换内容 - Json
	jsonBytes, errJ := json.Marshal(getResponse) //转换成JSON返回的是byte[]
	if errJ != nil {
		fmt.Println(errJ.Error())
	}
	fmt.Println(string(jsonBytes))

	// 打印转换内容 - Json 缩进
	jsonIndentBytes, errJI := json.MarshalIndent(getResponse, "", "\t") //转换成JSON返回的是byte[]
	if errJI != nil {
		fmt.Println(errJI.Error())
	}
	fmt.Println(string(jsonIndentBytes))
}

func post() {
	ro := &grequests.RequestOptions{
		JSON:   []byte(`{"One":"Two"}`),
		IsAjax: true,
	}

	resp, err := grequests.Post("http://httpbin.org/post", ro)

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	if resp.Error != nil {
		log.Fatalln("Unable to make request", resp.Error)
	}

	if resp.Ok != true {
		log.Fatalln("Request did not return OK")
	}

	type PostJSONResponse struct {
		Args    struct{} `json:"args"`
		Data    string   `json:"data"`
		Files   struct{} `json:"files"`
		Form    struct{} `json:"form"`
		Headers struct {
			AcceptEncoding string `json:"Accept-Encoding"`
			ContentLength  string `json:"Content-Length"`
			ContentType    string `json:"Content-Type"`
			Host           string `json:"Host"`
			UserAgent      string `json:"User-Agent"`
			XRequestedWith string `json:"X-Requested-With"`
		} `json:"headers"`
		JSON struct {
			One string `json:"One"`
		} `json:"json"`
		Origin string `json:"origin"`
		URL    string `json:"url"`
	}

	postJSONResponse := &PostJSONResponse{}
	if err := resp.JSON(postJSONResponse); err != nil {
		fmt.Println(err.Error())
	}
	// 打印构造结构
	fmt.Println(postJSONResponse)
	fmt.Println(postJSONResponse.URL)

	// 打印转换内容 - Json
	jsonBytes, errJ := json.Marshal(postJSONResponse) //转换成JSON返回的是byte[]
	if errJ != nil {
		fmt.Println(errJ.Error())
	}
	fmt.Println(string(jsonBytes))

	// 打印转换内容 - Json 缩进
	jsonIndentBytes, errJI := json.MarshalIndent(postJSONResponse, "", "\t") //转换成JSON返回的是byte[]
	if errJI != nil {
		fmt.Println(errJI.Error())
	}
	fmt.Println(string(jsonIndentBytes))
}

func upload() {
	fd, err := grequests.FileUploadFromDisk("downloads/test.md")

	if err != nil {
		log.Fatalln("Unable to open file: ", err)
	}

	resp, _ := grequests.Post("http://httpbin.org/post",
		&grequests.RequestOptions{
			Files: fd,
			Data:  map[string]string{"One": "Two"},
		})

	if resp.Error != nil {
		log.Fatalln("Unable to make request", resp.Error)
	}

	if resp.Ok != true {
		log.Fatalln("Request did not return OK")
	}

	type PostFileUploadResponse struct {
		Args  struct{} `json:"args"`
		Data  string   `json:"data"`
		Files struct {
			File string `json:"file"`
		} `json:"files"`
		Form struct {
			One string `json:"one"`
		} `json:"form"`
		Headers struct {
			AcceptEncoding string `json:"Accept-Encoding"`
			ContentLength  string `json:"Content-Length"`
			ContentType    string `json:"Content-Type"`
			Host           string `json:"Host"`
			UserAgent      string `json:"User-Agent"`
		} `json:"headers"`
		JSON   interface{} `json:"json"`
		Origin string      `json:"origin"`
		URL    string      `json:"url"`
	}

	postFileUploadResponse := &PostFileUploadResponse{}

	if err := resp.JSON(postFileUploadResponse); err != nil {
		log.Fatalln("Unable to coerce to JSON", err)
	}
	// 打印转换内容 - Json 缩进
	jsonIndentBytes, errJI := json.MarshalIndent(postFileUploadResponse, "", "\t") //转换成JSON返回的是byte[]
	if errJI != nil {
		fmt.Println(errJI.Error())
	}
	fmt.Println(string(jsonIndentBytes))
}

func main() {
	// 会话请求
	//getSession()

	// 基本请求
	//get()
	//post()
	upload()
}
