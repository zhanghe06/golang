package main

import (
	"fmt"
	"path"
	"runtime"
)

func getProjectAbPathByCaller() {
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		filePath := path.Dir(filename)
		abPath := path.Dir(filePath)
		fmt.Println(abPath)
	}
	return
}

func main()  {
	getProjectAbPathByCaller()
}
