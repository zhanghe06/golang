package main

/**
依赖github.com/alibabacloud-go/imageseg-20191230/v2
建议使用go mod tidy安装依赖
*/

import (
	"fmt"
	"os"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	imageaudit20191230 "github.com/alibabacloud-go/imageaudit-20191230/v3/client"
	imageseg20191230 "github.com/alibabacloud-go/imageseg-20191230/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliConfig struct {
	//Endpoint string
	AccessKeyId     string
	AccessKeySecret string
}

// ScanImage
// 图片格式：PNG、JPG、JPEG、BMP、GIF、WEBP。 图片大小：小于3MB。 图片像素：分辨率大于256×256。
func ScanImage(conf AliConfig) {

	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: tea.String(conf.AccessKeyId),
		// 必填，您的 AccessKey Secret
		AccessKeySecret: tea.String(conf.AccessKeySecret),
	}
	// 访问的域名
	config.Endpoint = tea.String("imageaudit.cn-shanghai.aliyuncs.com")
	client, err := imageaudit20191230.NewClient(config)
	if err != nil {
		panic(err)
	}
	// 场景一，使用本地文件
	//file, err := os.Open("/tmp/SegmentBody.png")
	file, err := os.Open("examples/images/正常.png")
	if err != nil {
		fmt.Println("can not open file", err)
		panic(err)
	}
	scene := make([]*string, 0)
	scene = append(scene, tea.String("porn"), tea.String("terrorism")) // 鉴黄、反恐

	task := make([]*imageaudit20191230.ScanImageAdvanceRequestTask, 0)
	task = append(task, &imageaudit20191230.ScanImageAdvanceRequestTask{
		ImageURLObject: file,
	})
	scanImageAdvanceRequest := &imageaudit20191230.ScanImageAdvanceRequest{
		Scene: scene,
		Task:  task,
	}

	runtime := &util.RuntimeOptions{}

	scanImageResponse, err := client.ScanImageAdvance(scanImageAdvanceRequest, runtime)
	if err != nil {
		// 获取整体报错信息
		switch interface{}(err).(type) {
		// 接口错误
		case *tea.SDKError:
			fmt.Println(*err.(*tea.SDKError).Code)
			fmt.Println(*err.(*tea.SDKError).Message)
			return
		// 其他异常
		default:
			fmt.Println(err.Error())
			return
		}
	} else {
		// 获取整体结果
		//fmt.Println(scanImageResponse)
		fmt.Println(*scanImageResponse.Body.Data)
	}
	if len(scanImageResponse.Body.Data.Results) == 0 {
		fmt.Println("result is empty")
		return
	}
	for i, v := range scanImageResponse.Body.Data.Results[0].SubResults {
		fmt.Println(i, *v.Suggestion == "block")
	}
}

func SegmentBodySame(conf AliConfig) {
	config := &openapi.Config{
		// "YOUR_ACCESS_KEY_ID", "YOUR_ACCESS_KEY_SECRET" 的生成请参考https://help.aliyun.com/document_detail/175144.html
		// 如果您是用的子账号AccessKey，还需要为子账号授予权限AliyunVIAPIFullAccess，请参考https://help.aliyun.com/document_detail/145025.html
		// 您的 AccessKey ID
		//AccessKeyId: tea.String("YOUR_ACCESS_KEY_ID"),
		AccessKeyId: tea.String(conf.AccessKeyId),
		// 您的 AccessKey Secret
		//AccessKeySecret: tea.String("YOUR_ACCESS_KEY_SECRET"),
		AccessKeySecret: tea.String(conf.AccessKeySecret),
	}
	// 访问的域名
	config.Endpoint = tea.String("imageseg.cn-shanghai.aliyuncs.com")
	client, err := imageseg20191230.NewClient(config)
	if err != nil {
		panic(err)
	}
	segmentBodyRequest := &imageseg20191230.SegmentBodyRequest{
		//ImageURL: tea.String("http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentBody/SegmentBody1.png"),
		ImageURL: tea.String("https://up.hczm1.com/edpic_source/64/d8/df/64d8dfe29262c2f4d9a26d59c9654c4b.jpg"),
	}
	runtime := &util.RuntimeOptions{}
	segmentBodyResponse, err := client.SegmentBodyWithOptions(segmentBodyRequest, runtime)
	if err != nil {
		// 获取整体报错信息
		switch interface{}(err).(type) {
		// 接口错误
		case *tea.SDKError:
			fmt.Println(*err.(*tea.SDKError).Code)
			fmt.Println(*err.(*tea.SDKError).Message)
			return
		// 其他异常
		default:
			fmt.Println(err.Error())
			return
		}
	} else {
		// 获取整体结果
		fmt.Println(segmentBodyResponse)
		fmt.Println(*segmentBodyResponse.Body.Data.ImageURL)
	}
}

func SegmentBodyFile(conf AliConfig) {
	config := &openapi.Config{
		// "YOUR_ACCESS_KEY_ID", "YOUR_ACCESS_KEY_SECRET" 的生成请参考https://help.aliyun.com/document_detail/175144.html
		// 如果您是用的子账号AccessKey，还需要为子账号授予权限AliyunVIAPIFullAccess，请参考https://help.aliyun.com/document_detail/145025.html
		// 您的 AccessKey ID
		//AccessKeyId: tea.String("YOUR_ACCESS_KEY_ID"),
		AccessKeyId: tea.String(conf.AccessKeyId),
		// 您的 AccessKey Secret
		//AccessKeySecret: tea.String("YOUR_ACCESS_KEY_SECRET"),
		AccessKeySecret: tea.String(conf.AccessKeySecret),
	}
	// 访问的域名
	config.Endpoint = tea.String("imageseg.cn-shanghai.aliyuncs.com")
	client, err := imageseg20191230.NewClient(config)
	if err != nil {
		panic(err)
	}
	// 场景一，使用本地文件
	//file, err := os.Open("/tmp/SegmentBody.png")
	file, err := os.Open("examples/images/非人.png")
	if err != nil {
		fmt.Println("can not open file", err)
		panic(err)
	}
	segmentBodyAdvanceRequest := &imageseg20191230.SegmentBodyAdvanceRequest{
		ImageURLObject: file,
	}
	// 场景二，使用任意可访问的url
	//httpClient := http.Client{}
	//file, _ := httpClient.Get("https://viapi-test-bj.oss-cn-beijing.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentBody/SegmentBody9.png")
	//segmentBodyAdvanceRequest := &imageseg20191230.SegmentBodyAdvanceRequest{
	//	ImageURLObject: file.Body,
	//}
	runtime := &util.RuntimeOptions{}
	segmentBodyAdvanceResponse, err := client.SegmentBodyAdvance(segmentBodyAdvanceRequest, runtime)
	if err != nil {
		// 获取整体报错信息
		switch interface{}(err).(type) {
		// 接口错误
		case *tea.SDKError:
			fmt.Println(*err.(*tea.SDKError).Code)
			fmt.Println(*err.(*tea.SDKError).Message)
			return
		// 其他异常
		default:
			fmt.Println(err.Error())
			return
		}
	} else {
		// 获取整体结果
		fmt.Println(segmentBodyAdvanceResponse)
		fmt.Println(*segmentBodyAdvanceResponse.Body.Data.ImageURL)
	}
}

func OSS() {
	client, err := oss.New(
		"Endpoint",
		"AccessKeyId",
		"AccessKeySecret",
	)
	if err != nil {
		// HandleError(err)
	}

	bucket, err := client.Bucket("my-bucket")
	if err != nil {
		// HandleError(err)
	}

	err = bucket.PutObjectFromFile("my-object", "LocalFile")
	if err != nil {
		// HandleError(err)
	}
}

func main() {
	var aliConfig = AliConfig{
		AccessKeyId:     os.Getenv("ACCESS_KEY_ID"),
		AccessKeySecret: os.Getenv("ACCESS_KEY_SECRET"),
	}

	//filePath := "examples/images/涉黄.png"
	//filePath := "examples/images/涉政.png"

	//ScanImage(aliConfig) // 图片审查
	//SegmentBodySame(aliConfig) // 人体分割（同域）
	SegmentBodyFile(aliConfig) // 人体分割（文件）
}
