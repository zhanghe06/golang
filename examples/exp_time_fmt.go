package main

import (
	"fmt"
	"time"
)

/*
2019-07-03 14:01:12.371795 +0800 CST m=+0.000158910
1562133672
2019-07-03 14:01:12
*/
func main() {
	now := time.Now()
	fmt.Println(now)

	nowUnix := time.Now().Unix()
	fmt.Println(nowUnix)

	// 当前格式化时间
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// 这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5

	// 时间偏移
	//am, _ := time.ParseDuration("5m")
	//sm, _ := time.ParseDuration("-5s")
	//
	//fmt.Println(time.Now().Add(am).Add(sm))

	am, _ := time.ParseDuration(fmt.Sprintf("%dm", 5))
	sm, _ := time.ParseDuration(fmt.Sprintf("-%ds", 5))
	fmt.Println(time.Now().Add(am).Add(sm))

	// UTC 时间
	fmt.Println(time.Now().Format(time.RFC3339))                 // 2022-08-05T09:45:14+08:00
	fmt.Println(time.Now().UTC().Format(time.RFC3339))           // 2022-08-05T01:45:14Z
	fmt.Println(time.Now().UTC().Format("2006-01-02T15:04:05Z")) // 2022-08-05T01:45:14Z

	// 时间戳表示唯一且具体的时间点，没有时区概念
	fmt.Println(time.Now().Unix())       // 1659663914
	fmt.Println(time.Now().UTC().Unix()) // 1659663914

	// 当前日期
	fmt.Println(time.Now().Format("2006-01-02"))

	fmt.Println(time.Now().Format("20060102150405"))

}
