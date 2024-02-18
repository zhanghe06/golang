package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// 针对大文件转储场景
// 对比 ioutil.ReadAll 和 io.Copy 内存占用情况

func testIoReadAll(fileSrcPath, fileDstPath string) {
	fileSrc, err := os.Open(fileSrcPath)
	if err != nil {
		log.Fatal(err)
	}
	defer fileSrc.Close()
	// ioutil.ReadAll 将把内容全部加载到内存中
	content, err := ioutil.ReadAll(fileSrc)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(fileDstPath, content, 0666)
	if err != nil {
		log.Fatal(err)
	}
	// 测试文件：87.2M；内存占用：206.2MB
	time.Sleep(time.Hour)
}

func testIoCopy(fileSrcPath, fileDstPath string) {
	fileSrc, err := os.Open(fileSrcPath)
	if err != nil {
		log.Fatal(err)
	}
	defer fileSrc.Close()
	//content := bytes.NewBuffer([]byte{})
	//var content *bytes.Buffer
	fileDst, err := os.Create(fileDstPath)
	if err != nil {
		log.Fatal(err)
	}
	defer fileDst.Close()
	// io.Copy 使用固定的 32KB 缓冲区从读取器复制到写入器，直到 EOF。因此，无论源有多大，总是只使用 32KB 将其复制到目标。
	written, err := io.Copy(fileDst, fileSrc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("written: %d\n", written)
	// 测试文件：87.2M；内存占用：948KB
	time.Sleep(time.Hour)
}

func main() {
	fileSrcPath := "/Users/zhanghe/Downloads/阅爱团队推荐书籍.zip"
	//fileDstPath := "/Users/zhanghe/Downloads/阅爱团队推荐书籍-1.zip"
	fileDstPath := "/Users/zhanghe/Downloads/阅爱团队推荐书籍-2.zip"
	//testIoReadAll(fileSrcPath, fileDstPath)
	testIoCopy(fileSrcPath, fileDstPath)
}
