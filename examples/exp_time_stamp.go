package main

import (
	"fmt"
	"time"
)

// GetCurrentTimestamp 获取当天的时间范围
//Time类型 2022-05-19 00:00:00 +0800 CST 2022-05-19 23:59:59 +0800 CST
func GetCurrentTimestamp() (beginTime, endTime time.Time) {
	t := time.Now()
	timeStr := t.Format("2006-01-02")
	beginTime, _ = time.ParseInLocation("2006-01-02", timeStr, time.Local)
	endTimeTmp := beginTime.Unix() + 86399
	endTime = time.Unix(endTimeTmp, 0)
	return beginTime, endTime
}


func GetTodayUnix() (begin, end int64) {
	t := time.Now()
	timeStr := t.Format("2006-01-02")
	todayBeginTime, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	todayEndTimeTmp := todayBeginTime.Unix() + 86399
	todayEndTime := time.Unix(todayEndTimeTmp, 0)

	begin = todayBeginTime.Unix()
	end = todayEndTime.Unix()
	return
}


func GetYesterdayUnix() (begin, end int64) {
	t := time.Now()
	timeStr := t.Format("2006-01-02")
	todayBeginTime, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	todayEndTimeTmp := todayBeginTime.Unix() + 86399
	todayEndTime := time.Unix(todayEndTimeTmp, 0)

	begin = todayBeginTime.Unix() - 24*3600
	end = todayEndTime.Unix() - 24*3600
	return
}

func GetBeforeYesterdayUnix() (begin, end int64) {
	t := time.Now()
	timeStr := t.Format("2006-01-02")
	todayBeginTime, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	todayEndTimeTmp := todayBeginTime.Unix() + 86399
	todayEndTime := time.Unix(todayEndTimeTmp, 0)

	begin = todayBeginTime.Unix() - 24*3600*2
	end = todayEndTime.Unix() - 24*3600*2
	return
}



func main()  {
	fmt.Println(GetCurrentTimestamp())
	fmt.Println(GetTodayUnix())
	fmt.Println(GetYesterdayUnix())
	fmt.Println(GetBeforeYesterdayUnix())
	fmt.Println(time.Now().Unix()) // 时间戳，单位：秒

	fmt.Println(time.Now().Unix() + time.Duration(1000000).Milliseconds())
}