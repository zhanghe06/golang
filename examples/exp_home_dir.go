package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	u, err := user.Current()
	if err != nil {
		fmt.Println("读主目录错误：", err)
		return
	}
	fmt.Println("Home dir:", u.HomeDir)

	err = os.MkdirAll(filepath.Join(u.HomeDir, "pornhub"), os.ModePerm)
	if err != nil {
		fmt.Println("创建目录错误：", err)
		return
	}
}
