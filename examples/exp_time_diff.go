package main

import (
	"fmt"
)

func timeDiffUnix(bef, now int64)  {
	fmt.Println((now-bef)/3600/24)
}

func main()  {
	timeDiffUnix(1664726400, 1665331199)
	timeDiffUnix(1664726400, 1665331200)
}