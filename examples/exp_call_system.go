package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("ls", "-alh")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("stdout:", stdout.String())
	fmt.Println("stderr:", stderr.String())
}
