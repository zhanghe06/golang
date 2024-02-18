package main

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"log"
)

const (
	//EndpointChartInstance = "http://172.16.71.154:31011/kubernetes/clustertiller"
	EndpointChartInstance = "http://172.16.25.41:31011/kubernetes/clustertiller"
	UrlChartInstanceGet   = EndpointChartInstance + "/getChartInstance"
	UrlChartInstanceDel   = EndpointChartInstance + "/deleteChartInstance"
	UrlChartInstanceAdd   = EndpointChartInstance + "/createChartInstance"
)

// 详情
type ChartInstanceGetRequest struct {
	ClusterID   string `json:"clusterID"`
	NamespaceID string `json:"namespaceID"`
	ReleaseName string `json:"releaseName"`
}

type ChartInstanceGetResponse struct {
	Data      map[string]interface{} `json:"data"`
	RequestId string                 `json:"requestId"`
}

// 删除
type ChartInstanceDelRequest struct {
	ClusterID   string `json:"clusterID"`
	NamespaceID string `json:"namespaceID"`
	ReleaseName string `json:"releaseName"`
}

type ChartInstanceDelResponse struct {
	Data      map[string]interface{} `json:"data"`
	RequestId string                 `json:"requestId"`
}

// 创建
type ChartInstanceAddRequest struct {
	ClusterID   string `json:"clusterID"`
	NamespaceID string `json:"namespaceID"`
	ReleaseName string `json:"releaseName"`
	ChartName   string `json:"chartName"`
	RepoUrl     string `json:"repoUrl"`
	Version     string `json:"version"`
	Values      string `json:"values"`
}

type ChartInstanceAddResponse struct {
	Data struct {
		Saved bool `json:"saved"`
	} `json:"data"`
	RequestId string `json:"requestId"`
}

func AddChartInstance(chartInstanceAddRequest ChartInstanceAddRequest) (*ChartInstanceAddResponse, error) {
	jsonChartInstanceAddRequest, err := json.Marshal(chartInstanceAddRequest) //转换成JSON返回的是byte[]
	if err != nil {
		return nil, fmt.Errorf("unable to coerce to JSON: %v", err)
		//fmt.Println(err.Error())
	}

	ro := &grequests.RequestOptions{
		JSON: jsonChartInstanceAddRequest,
	}

	resp, err := grequests.Post(UrlChartInstanceAdd, ro)

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

	chartInstanceAddResponse := &ChartInstanceAddResponse{}
	if err := resp.JSON(chartInstanceAddResponse); err != nil {
		return nil, fmt.Errorf("unable to coerce to JSON: %v", err)
		//fmt.Println(err.Error())
	}
	return chartInstanceAddResponse, nil
}

func GetChartInstance(chartInstanceGetRequest ChartInstanceGetRequest) (*ChartInstanceGetResponse, error) {

	jsonChartInstanceGetRequest, err := json.Marshal(chartInstanceGetRequest) //转换成JSON返回的是byte[]
	if err != nil {
		return nil, fmt.Errorf("unable to coerce to JSON: %v", err)
		//fmt.Println(err.Error())
	}

	ro := &grequests.RequestOptions{
		JSON: jsonChartInstanceGetRequest,
	}

	resp, err := grequests.Post(UrlChartInstanceGet, ro)

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

	ChartInstanceGetResponse := &ChartInstanceGetResponse{}
	if err := resp.JSON(ChartInstanceGetResponse); err != nil {
		return nil, fmt.Errorf("unable to coerce to JSON: %v", err)
		//fmt.Println(err.Error())
	}
	return ChartInstanceGetResponse, nil
}

func DelChartInstance(chartInstanceDelRequest ChartInstanceDelRequest) (*ChartInstanceDelResponse, error) {

	jsonChartInstanceDelRequest, err := json.Marshal(chartInstanceDelRequest) //转换成JSON返回的是byte[]
	if err != nil {
		return nil, fmt.Errorf("unable to coerce to JSON: %v", err)
		//fmt.Println(err.Error())
	}

	ro := &grequests.RequestOptions{
		JSON: jsonChartInstanceDelRequest,
	}

	resp, err := grequests.Post(UrlChartInstanceDel, ro)

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

	ChartInstanceDelResponse := &ChartInstanceDelResponse{}
	if err := resp.JSON(ChartInstanceDelResponse); err != nil {
		return nil, fmt.Errorf("unable to coerce to JSON: %v", err)
		//fmt.Println(err.Error())
	}
	return ChartInstanceDelResponse, nil
}

func testChartInstanceAdd() {
	// 测试创建
	chartInstanceAddRequest := ChartInstanceAddRequest{
		ClusterID:   "337bdedf-b468-4964-abe3-ca7c4d53fe2a",
		NamespaceID: "default",
		ReleaseName: "my-nginx-chart-release",
		ChartName:   "nginx",
		RepoUrl:     "https://charts.bitnami.com/bitnami",
		Version:     "3.1.0",
		Values:      "",
	}
	response, err := AddChartInstance(chartInstanceAddRequest)
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

func testChartInstanceGet() {
	// 测试详情
	chartInstanceGetRequest := ChartInstanceGetRequest{
		//ClusterID:   "337bdedf-b468-4964-abe3-ca7c4d53fe2a",
		//NamespaceID: "kubeapps",
		//ReleaseName: "kubeapps",
		ClusterID:   "panfengyun_cluster_2",
		NamespaceID: "default",
		ReleaseName: "garish-sloth",
	}

	response, err := GetChartInstance(chartInstanceGetRequest)
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

func testChartInstanceDel() {
	// 测试删除
	chartInstanceDelRequest := ChartInstanceDelRequest{
		ClusterID:   "337bdedf-b468-4964-abe3-ca7c4d53fe2a",
		NamespaceID: "default",
		ReleaseName: "my-nginx-chart-release",
	}

	response, err := DelChartInstance(chartInstanceDelRequest)
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
	// 测试创建
	//testChartInstanceAdd()

	// 测试详情
	testChartInstanceGet()

	// 测试删除
	//testChartInstanceDel()
}
