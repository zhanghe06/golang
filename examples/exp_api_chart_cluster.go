package main

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"gopkg.in/yaml.v2"
	"log"
)

const (
	//EndpointClusterChart     = "http://172.16.71.154:31011/kubernetes/clusterchart"
	EndpointClusterChart     = "http://172.16.25.41:31011/kubernetes/clusterchart"
	UrlGetClusterChart       = EndpointClusterChart + "/getClusterChart"
	UrlGetClusterChartReadme = EndpointClusterChart + "/getClusterChartReadme"
	UrlGetClusterChartValues = EndpointClusterChart + "/getClusterChartValues"
)

// ClusterChart
type GetClusterChartRequest struct {
	ClusterID    string `json:"clusterID"`
	ChartName    string `json:"chartName"`
	ChartVersion string `json:"chartVersion"`
}

type GetClusterChartResponse struct {
	Data      map[string]interface{} `json:"data"`
	RequestId string                 `json:"requestId"`
}

// ClusterChartReadme
type GetClusterChartReadmeRequest struct {
	ClusterID    string `json:"clusterID"`
	ChartName    string `json:"chartName"`
	ChartVersion string `json:"chartVersion"`
}

type GetClusterChartReadmeResponse struct {
	Data      string `json:"data"`
	RequestId string `json:"requestId"`
}

// ClusterChartValues
type GetClusterChartValuesRequest struct {
	ClusterID    string `json:"clusterID"`
	ChartName    string `json:"chartName"`
	ChartVersion string `json:"chartVersion"`
}

type GetClusterChartValuesResponse struct {
	Data      string `json:"data"`
	RequestId string `json:"requestId"`
}

func GetClusterChart(getClusterChartRequest GetClusterChartRequest) (*GetClusterChartResponse, error) {

	jsonGetClusterChartRequest, err := json.Marshal(getClusterChartRequest) //转换成JSON返回的是byte[]
	if err != nil {
		return nil, fmt.Errorf("unable to coerce to JSON: %v", err)
		//fmt.Println(err.Error())
	}

	ro := &grequests.RequestOptions{
		JSON: jsonGetClusterChartRequest,
	}

	resp, err := grequests.Post(UrlGetClusterChart, ro)

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

	getClusterChartResponse := &GetClusterChartResponse{}
	if err := resp.JSON(getClusterChartResponse); err != nil {
		return nil, fmt.Errorf("unable to coerce to JSON: %v", err)
		//fmt.Println(err.Error())
	}
	return getClusterChartResponse, nil
}

func GetClusterChartReadme(getClusterChartReadmeRequest GetClusterChartReadmeRequest) (*GetClusterChartReadmeResponse, error) {

	jsonGetClusterChartReadmeRequest, err := json.Marshal(getClusterChartReadmeRequest) //转换成JSON返回的是byte[]
	if err != nil {
		return nil, fmt.Errorf("unable to coerce to JSON: %v", err)
		//fmt.Println(err.Error())
	}

	ro := &grequests.RequestOptions{
		JSON: jsonGetClusterChartReadmeRequest,
	}

	resp, err := grequests.Post(UrlGetClusterChartReadme, ro)

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

	getClusterChartReadmeResponse := &GetClusterChartReadmeResponse{}
	if err := resp.JSON(getClusterChartReadmeResponse); err != nil {
		return nil, fmt.Errorf("unable to coerce to JSON: %v", err)
		//fmt.Println(err.Error())
	}
	return getClusterChartReadmeResponse, nil
}

func GetClusterChartValues(getClusterChartValuesRequest GetClusterChartValuesRequest) (*GetClusterChartValuesResponse, error) {

	jsonGetClusterChartValuesRequest, err := json.Marshal(getClusterChartValuesRequest) //转换成JSON返回的是byte[]
	if err != nil {
		return nil, fmt.Errorf("unable to coerce to JSON: %v", err)
		//fmt.Println(err.Error())
	}

	ro := &grequests.RequestOptions{
		JSON: jsonGetClusterChartValuesRequest,
	}

	resp, err := grequests.Post(UrlGetClusterChartValues, ro)

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

	getClusterChartValuesResponse := &GetClusterChartValuesResponse{}
	if err := resp.JSON(getClusterChartValuesResponse); err != nil {
		return nil, fmt.Errorf("unable to coerce to JSON: %v", err)
		//fmt.Println(err.Error())
	}
	return getClusterChartValuesResponse, nil
}

func testGetClusterChart() {
	// 测试ClusterChart
	getClusterChartRequest := GetClusterChartRequest{
		//ClusterID:    "337bdedf-b468-4964-abe3-ca7c4d53fe2a",
		//ChartName:    "bitnami/redis",
		//ChartVersion: "7.1.0",
		ClusterID:    "panfengyun_cluster_2",
		ChartName:    "aliyun/mariadb",
		ChartVersion: "2.1.6",
	}

	response, err := GetClusterChart(getClusterChartRequest)
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

func testGetClusterChartReadme() {
	// 测试ClusterChartReadme
	getClusterChartReadmeRequest := GetClusterChartReadmeRequest{
		ClusterID:    "337bdedf-b468-4964-abe3-ca7c4d53fe2a",
		ChartName:    "bitnami/redis",
		ChartVersion: "7.1.0",
	}

	response, err := GetClusterChartReadme(getClusterChartReadmeRequest)
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

func testGetClusterChartValues() {
	// 测试ClusterChartValues
	getClusterChartValuesRequest := GetClusterChartValuesRequest{
		//ClusterID:    "337bdedf-b468-4964-abe3-ca7c4d53fe2a",
		ClusterID:    "panfengyun_cluster_2",
		ChartName:    "aliyun/mariadb",
		ChartVersion: "2.1.6",
	}

	response, err := GetClusterChartValues(getClusterChartValuesRequest)
	if err != nil {
		log.Fatalln(err)
	}

	// yaml 转 map
	mapObj := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(response.Data), &mapObj)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(mapObj)
	// 读取嵌套map
	//fmt.Println(mapObj["image"].(map[interface{}]interface{})["repository"])

	// map 转 yaml
	yamlBytes, err := yaml.Marshal(&mapObj)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(yamlBytes))

	// 打印转换内容 - Json 缩进
	jsonIndentBytes, errJI := json.MarshalIndent(response, "", "\t") //转换成JSON返回的是byte[]
	if errJI != nil {
		fmt.Println(errJI.Error())
	}
	fmt.Println(string(jsonIndentBytes))
}

func main() {
	// 测试ClusterChart
	testGetClusterChart()

	// 测试ClusterChartReadme
	//testGetClusterChartReadme()

	// 测试ClusterChartValues
	//testGetClusterChartValues()
}
