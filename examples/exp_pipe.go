package main

import (
	"errors"
	"fmt"
	"io"
	"time"
)

func main() {
	pipeReader, pipeWriter := io.Pipe()
	go pipeWrite(pipeWriter)
	go pipeRead(pipeReader)
	time.Sleep(20 * time.Second)
}

func pipeRead(reader *io.PipeReader) {
	buf := make([]byte, 3) // 每个汉字3字节
	fmt.Println("开始接收")
	for {
		n, err := reader.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		// 注意这里buf[:n], 不能直接用buf
		fmt.Printf("收到字节: %d\t内容: %s\n", n, buf[:n])
		// 每次循环清空buf，打印出来可能会存在空字节
		// fmt.Printf("收到字节: %d\t内容: %s\n", n, buf)
		// buf = make([]byte, 3)
	}
}

func pipeWrite(write *io.PipeWriter) {
	data := []byte("管道测试12345678")
	for i := 0; i < 3; i++ { //写入次数 i
		time.Sleep(1 * time.Second)
		n, err := write.Write(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("写入字节: %d\t内容: %s\n", n, data)
	}

	_ = write.CloseWithError(errors.New("关闭写入"))
}
