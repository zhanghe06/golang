package main

import (
	"fmt"
	"sort"
	"time"
)

func GetUnix(n int64) (begin, end int64) {
	t := time.Now()
	timeStr := t.Format("2006-01-02")
	todayBeginTime, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	todayEndTimeTmp := todayBeginTime.Unix() + 86399
	todayEndTime := time.Unix(todayEndTimeTmp, 0)

	begin = todayBeginTime.Unix() - 24*3600*n
	end = todayEndTime.Unix() - 24*3600*n
	return
}

func main()  {
	var startTime int64
	var endTime int64
	startTime = 1664726400
	endTime = 1664812799
	var timeRangeSlice [][]int64
	var n int64
	for {
		begin, end := GetUnix(n)
		timeRangeSlice = append(timeRangeSlice, []int64{begin, end})
		if begin == startTime && end == endTime {
			break
		}
		n++
	}

	// 倒序排列
	sort.SliceStable(timeRangeSlice, func(i, j int) bool {
		return true
	})

	fmt.Println(timeRangeSlice)
}
