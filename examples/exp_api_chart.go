package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/levigross/grequests"
)

const (
	EndpointChart  = "http://172.16.71.154:31011/kubernetes/chart"
	UrlChartGet    = EndpointChart + "/getChart"
	UrlChartDel    = EndpointChart + "/deleteChart"
	UrlChartUpload = EndpointChart + "/uploadChart"
)

// ChartGetRequest 详情
type ChartGetRequest struct {
	ChartName    string `json:"chartName"`
	ChartVersion string `json:"chartVersion"`
}

type ChartGetResponse struct {
	Data      map[string]interface{} `json:"data"`
	RequestId string                 `json:"requestId"`
}

// ChartDelRequest 删除
type ChartDelRequest struct {
	ChartName    string `json:"chartName"`
	ChartVersion string `json:"chartVersion"`
}

type ChartDelResponse struct {
	Data      map[string]interface{} `json:"data"`
	RequestId string                 `json:"requestId"`
}

// ChartUploadResponse 上传
type ChartUploadResponse struct {
	Data struct {
		Saved bool `json:"saved"`
	} `json:"data"`
	RequestId string `json:"requestId"`
}

func UploadChart(filePath string) (*ChartUploadResponse, error) {
	// filePath "downloads/test.md"
	fd, err := grequests.FileUploadFromDisk(filePath)

	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
		//log.Fatalln("Unable to open file: ", err)
	}

	resp, _ := grequests.Post(UrlChartUpload,
		&grequests.RequestOptions{
			Files: fd,
		})

	if resp.Error != nil {
		return nil, fmt.Errorf("unable to make request: %v", resp.Error)
		//log.Fatalln("Unable to make request", resp.Error)
	}

	if resp.Ok != true {
		errorMsg := fmt.Errorf("request did not return OK")
		if resp.StatusCode == 404 {
			errorMsg = fmt.Errorf("not found")
		}
		if resp.StatusCode == 500 {
			errorMsg = fmt.Errorf("internal server error")
		}
		return nil, errorMsg
		//log.Fatalln("Request did not return OK")
	}

	chartUploadResponse := &ChartUploadResponse{}

	if err := resp.JSON(chartUploadResponse); err != nil {
		return nil, fmt.Errorf("unable to coerce to JSON: %v", err)
		//log.Fatalln("Unable to coerce to JSON", err)
	}

	return chartUploadResponse, nil
}

func GetChart(chartGetRequest ChartGetRequest) (*ChartGetResponse, error) {

	jsonChartGetRequest, err := json.Marshal(chartGetRequest) //转换成JSON返回的是byte[]
	if err != nil {
		return nil, fmt.Errorf("unable to coerce to JSON: %v", err)
		//fmt.Println(err.Error())
	}

	ro := &grequests.RequestOptions{
		JSON: jsonChartGetRequest,
	}

	resp, err := grequests.Post(UrlChartGet, ro)

	if err != nil {
		return nil, fmt.Errorf("unable to make request: %v", err)
		//log.Fatalln("Unable to make request: ", err)
	}

	if resp.Error != nil {
		return nil, fmt.Errorf("unable to make request: %v", resp.Error)
		//log.Fatalln("Unable to make request", resp.Error)
	}

	if resp.Ok != true {
		errorMsg := fmt.Errorf("request did not return OK")
		if resp.StatusCode == 404 {
			errorMsg = fmt.Errorf("not found")
		}
		if resp.StatusCode == 500 {
			errorMsg = fmt.Errorf("internal server error")
		}
		return nil, errorMsg
		//log.Fatalln("Request did not return OK")
	}

	chartGetResponse := &ChartGetResponse{}
	if err := resp.JSON(chartGetResponse); err != nil {
		return nil, fmt.Errorf("unable to coerce to JSON: %v", err)
		//fmt.Println(err.Error())
	}
	return chartGetResponse, nil
}

func DelChart(chartDelRequest ChartDelRequest) (*ChartDelResponse, error) {

	jsonChartDelRequest, err := json.Marshal(chartDelRequest) //转换成JSON返回的是byte[]
	if err != nil {
		return nil, fmt.Errorf("unable to coerce to JSON: %v", err)
		//fmt.Println(err.Error())
	}

	ro := &grequests.RequestOptions{
		JSON: jsonChartDelRequest,
	}

	resp, err := grequests.Post(UrlChartDel, ro)

	if err != nil {
		return nil, fmt.Errorf("unable to make request: %v", err)
		//log.Fatalln("Unable to make request: ", err)
	}

	if resp.Error != nil {
		return nil, fmt.Errorf("unable to make request: %v", resp.Error)
		//log.Fatalln("Unable to make request", resp.Error)
	}

	if resp.Ok != true {
		errorMsg := fmt.Errorf("request did not return OK")
		if resp.StatusCode == 404 {
			errorMsg = fmt.Errorf("not found")
		}
		if resp.StatusCode == 500 {
			errorMsg = fmt.Errorf("internal server error")
		}
		return nil, errorMsg
		//log.Fatalln("Request did not return OK")
	}

	chartDelResponse := &ChartDelResponse{}
	if err := resp.JSON(chartDelResponse); err != nil {
		return nil, fmt.Errorf("unable to coerce to JSON: %v", err)
		//fmt.Println(err.Error())
	}
	return chartDelResponse, nil
}

func testChartUpload() {
	// 测试上传
	filePath := "downloads/test.md"
	response, err := UploadChart(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	// 打印转换内容 - Json 缩进
	jsonIndentBytes, errJI := json.MarshalIndent(response, "", "\t") //转换成JSON返回的是byte[]
	if errJI != nil {
		fmt.Println(errJI.Error())
	}
	fmt.Println(string(jsonIndentBytes))
}

func testChartGet() {
	// 测试详情
	chartGetRequest := ChartGetRequest{
		ChartName:    "mongodb",
		ChartVersion: "4.2.1",
	}

	response, err := GetChart(chartGetRequest)
	if err != nil {
		log.Fatalln(err)
	}

	// 打印转换内容 - Json 缩进
	jsonIndentBytes, errJI := json.MarshalIndent(response, "", "\t") //转换成JSON返回的是byte[]
	if errJI != nil {
		fmt.Println(errJI.Error())
	}
	fmt.Println(string(jsonIndentBytes))
}

func testChartDel() {
	// 测试删除
	chartDelRequest := ChartDelRequest{
		ChartName:    "nymph",
		ChartVersion: "0.2",
	}

	response, err := DelChart(chartDelRequest)
	if err != nil {
		log.Fatalln(err)
	}

	// 打印转换内容 - Json 缩进
	jsonIndentBytes, errJI := json.MarshalIndent(response, "", "\t") //转换成JSON返回的是byte[]
	if errJI != nil {
		fmt.Println(errJI.Error())
	}
	fmt.Println(string(jsonIndentBytes))
}

func main() {
	// 测试上传
	testChartUpload()

	// 测试详情
	testChartGet()

	// 测试删除
	testChartDel()
}
