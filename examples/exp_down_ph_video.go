package main

// 名称：PornHub视频下载助手
// 作者：https://github.com/zhanghe06

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	videoHome string
	videoName string
	videoLink string
	httpProxy string
)

// ProgressBar ...
type ProgressBar struct {
	percent int64  // progress percentage
	cur     int64  // current progress
	total   int64  // total value for progress
	rate    string // the actual progress bar to be printed
	graph   string // the fill value for progress bar
}

func (bar *ProgressBar) NewOption(start, total int64) {
	bar.cur = start
	bar.total = total
	if bar.graph == "" {
		//bar.graph = "█"
		bar.graph = "="
	}
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i += 2 {
		bar.rate += bar.graph // initial progress position
	}
}

func (bar *ProgressBar) getPercent() int64 {
	return int64((float32(bar.cur) / float32(bar.total)) * 100)
}

func (bar *ProgressBar) Reset() {
	bar.percent = 0
	bar.cur = 0
	bar.rate = ""
}

func (bar *ProgressBar) Play(cur int64, title string) {
	bar.cur = cur
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.graph
	}

	fmt.Printf("\r%s[%-50s]%3d%% %8d/%d", title, bar.rate, bar.percent, bar.cur, bar.total)
}

func (bar *ProgressBar) Finish() {
	fmt.Println()
}

type PornHub struct {
	videoLink string
	videoHost string
	videoName string
	videoHome string
	videoPath string
	httpProxy string
	urlFormat string
	maxNumber int
}

func (ph *PornHub) NewOption(videoLink, videoHome, videoName, httpProxy string) {
	if videoLink == "" {
		log.Fatalln("视频链接为空：")
		return
	}

	ph.videoLink = videoLink

	// 存储目录
	ph.videoHome = videoHome
	if ph.videoHome == "" {
		u, err := user.Current()
		if err != nil {
			log.Fatalln("用户目录错误：", err)
			return
		}
		ph.videoHome = filepath.Join(u.HomeDir, "pornhub")
		err = os.MkdirAll(ph.videoHome, os.ModePerm)
		if err != nil {
			log.Fatalln("创建目录错误：", err)
			return
		}
	}

	ph.videoName = videoName
	if ph.videoName == "" {
		ph.videoName = fmt.Sprintf("video-%s.ts", time.Now().Format("20060102150405"))
	}

	ph.videoPath = filepath.Join(ph.videoHome, ph.videoName)

	ph.httpProxy = httpProxy
	ph.parseVideoLink()
}

func (ph *PornHub) parseVideoLink() {
	tsURL, _ := url.Parse(ph.videoLink)
	ph.videoHost = tsURL.Host

	re := regexp.MustCompile("seg-(.*?)-f1-v1-a1.ts")

	sub := re.FindSubmatch([]byte(ph.videoLink))

	if len(sub) < 2 {
		log.Fatalln("视频地址格式错误")
		return
	}
	ph.maxNumber, _ = strconv.Atoi(string(sub[1]))

	segFormat := "seg-%d-f1-v1-a1.ts"
	ph.urlFormat = strings.ReplaceAll(ph.videoLink, "%", "%%")
	ph.urlFormat = strings.ReplaceAll(ph.urlFormat, fmt.Sprintf(segFormat, ph.maxNumber), segFormat)

	return
}

// P站.ts文件编码规则
func (ph *PornHub) getCodeList() (res []int) {
	// 2023-07-26 旧规则已经失效
	//if ph.videoHost == "ev-h.phncdn.com" {
	//	if ph.maxNumber < 10 {
	//		return
	//	}
	//	group := ph.maxNumber / 10 // 分组
	//	count := ph.maxNumber % 10 // 余数
	//	for i := range make([]struct{}, group) {
	//		res = append(res, i+1)
	//		num := 10
	//		if i+1 == group {
	//			num = count + 1
	//		}
	//		for j := range make([]struct{}, num) {
	//			code, _ := strconv.Atoi(fmt.Sprintf("%d%d", i+1, j))
	//			res = append(res, code)
	//		}
	//	}
	//}
	//if ph.videoHost == "cv-h.phncdn.com" {
	//	for i := range make([]struct{}, ph.maxNumber) {
	//		res = append(res, i+1)
	//	}
	//}

	for i := range make([]struct{}, ph.maxNumber) {
		res = append(res, i+1)
	}
	return
}

func (ph *PornHub) downVideoTS(videoTSLink, videoTSPath string) {
	request, err := http.NewRequest("GET", videoTSLink, nil)
	if err != nil {
		return
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if ph.httpProxy != "" {
		proxy, _ := url.Parse(ph.httpProxy)
		tr.Proxy = http.ProxyURL(proxy)
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(request)
	if err != nil {
		log.Println("error:", err)
		return
	}
	defer func() {
		_ = res.Body.Close()
	}()

	var content []byte
	content, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("读取内容错误：", err)
		return
	}

	var file *os.File
	file, err = os.OpenFile(videoTSPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println("打开文件错误：", err)
		return
	}
	defer func() {
		_ = file.Close()
	}()

	write := bufio.NewWriter(file)
	_, err = write.Write(content)
	if err != nil {
		log.Println("写入文件错误：", err)
		return
	}

	err = write.Flush()
	if err != nil {
		log.Println("刷入文件错误：", err)
		return
	}
}

func (ph *PornHub) mergeVideoTS(videoTSPath string) {
	outFile, err := os.OpenFile(ph.videoPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println("创建文件错误：", err)
		return
	}
	defer func() {
		_ = outFile.Close()
	}()

	var oriFile *os.File
	oriFile, err = os.OpenFile(videoTSPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Println("打开文件错误：", err)
		return
	}
	defer func() {
		_ = oriFile.Close()
	}()

	var content []byte
	content, err = ioutil.ReadAll(oriFile)
	if err != nil {
		log.Println("读取文件错误：", err)
		return
	}
	_, err = outFile.Write(content)
	if err != nil {
		log.Println("合并文件错误：", err)
		return
	}
}

func (ph *PornHub) deleteVideoTS(videoTSPath string) {
	err := os.Remove(videoTSPath)
	if err != nil {
		log.Println("删除文件错误：", err)
		return
	}
}

func (ph *PornHub) runBash() {

	codeList := ph.getCodeList()
	videoTSList := make([]string, 0)

	for _, code := range codeList {
		videoTSUrl := fmt.Sprintf(ph.urlFormat, code)
		videoTSPath := filepath.Join(ph.videoHome, fmt.Sprintf("%d.ts", code))
		cmdDownload := fmt.Sprintf("wget -c -O %s %s\n", videoTSPath, videoTSUrl)
		fmt.Print(cmdDownload)

		videoTSList = append(videoTSList, videoTSPath)
	}
	cmdMerge := fmt.Sprintf("ffmpeg -i \"concat:%s\" -c copy %s", strings.Join(videoTSList, "|"), ph.videoPath)
	fmt.Println(cmdMerge)
}

func (ph *PornHub) runGo() {
	// 进度显示
	var progressBar ProgressBar

	codeList := ph.getCodeList()
	videoTSList := make([]string, 0)

	// 视频标识
	videoCode := uuid.New().String()

	// 下载
	progressBar.NewOption(0, int64(len(codeList)))
	for i := 0; i < len(codeList); i++ {
		progressBar.Play(int64(i), "下载(1/3)")

		code := codeList[i]
		videoTSLink := fmt.Sprintf(ph.urlFormat, code)
		videoTSPath := filepath.Join(ph.videoHome, videoCode, fmt.Sprintf("%d.ts", code))
		ph.downVideoTS(videoTSLink, videoTSPath)
		videoTSList = append(videoTSList, fmt.Sprintf("%d.ts", code))

		progressBar.Play(int64(i+1), "下载(1/3)")
	}
	// 合并
	progressBar.Reset()
	progressBar.NewOption(0, int64(len(codeList)))
	for i := 0; i < len(codeList); i++ {
		progressBar.Play(int64(i), "合并(2/3)")

		code := codeList[i]
		videoTSPath := filepath.Join(ph.videoHome, videoCode, fmt.Sprintf("%d.ts", code))
		ph.mergeVideoTS(videoTSPath)

		progressBar.Play(int64(i+1), "合并(2/3)")
	}
	// 清理
	progressBar.Reset()
	progressBar.NewOption(0, int64(len(codeList)))
	for i := 0; i < len(codeList); i++ {
		progressBar.Play(int64(i), "清理(3/3)")

		code := codeList[i]
		videoTSPath := filepath.Join(ph.videoHome, fmt.Sprintf("%d.ts", code))
		ph.deleteVideoTS(videoTSPath)

		progressBar.Play(int64(i+1), "清理(3/3)")
	}

	// 清理进度
	progressBar.Finish()
}

func main() {
	// 编译：
	// - GOOS：目标平台的操作系统（darwin、freebsd、linux、windows）
	// - GOARCH：目标平台的体系架构（386、amd64、arm）
	// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fuck_ph_amd64 exp_down_ph_video.go
	// CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o fuck_ph_arm exp_down_ph_video.go
	// CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o fuck_ph.exe exp_down_ph_video.go

	// go build -o fuck_ph exp_down_ph_video.go

	// 示例：
	// 将视频拉到最后，获取最后一个ts链接地址
	// "https://cv-h.phncdn.com/hls/videos/202301/21/423821271/,1080P_4000K,720P_4000K,480P_2000K,240P_1000K,_423821271.mp4.urlset/seg-273-f1-v1-a1.ts?nUqz9qnHZqwrp5ukxOXu0X39exOogj0s9KeoeUE32xKHknVfbtUMCo5osJdrQUmWFarXleIFcuDzLY2mt0uEwHOqZRuZGXjGr27p8cQ8i-pVCpz4Qeaw_YrkqDyDNCAt6O7gV6mrzfVHZCRD7qKUqHFE7H_pY2CMCNCXMlBKXY48mhMjvilxwccb2VyQ0lXHfhnLnw8-"

	// 执行：
	// ./fuck_ph -video_link="https://cv-h.phncdn.com/xxx" -http_proxy=http://127.0.0.1:1087

	// 参数解析
	flag.StringVar(&videoLink, "video_link", "", "请输入最后一段视频链接 .ts格式 需要加上双引号")
	flag.StringVar(&videoName, "video_name", "", "请输入视频文件名称（可选）.ts格式")
	flag.StringVar(&videoHome, "video_home", "", "请输入视频存储目录（可选）默认存储在用户目录的pornhub文件夹下")
	flag.StringVar(&httpProxy, "http_proxy", "", "请输入HTTP代理（可选）")
	flag.Parse()

	ph := PornHub{}
	ph.NewOption(videoLink, videoHome, videoName, httpProxy)

	ph.runGo()
}
