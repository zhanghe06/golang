package main

import (
	"log"
)

// bucket token算法
type ConnLimiter struct {
	concurrentConn int      //连接总数
	bucket         chan int //当前连接数
}

// 设置bucket token数量
func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,                 //设置连接总数
		bucket:         make(chan int, cc), //设置channel的长度
	}
}

// 获取一个连接
func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("连接数已满:%d", cl.concurrentConn)
		return false
	}
	//向channel中添加一个
	cl.bucket <- 1
	log.Printf("增加一个连接")
	//time.Sleep(time.Duration(10)*time.Second)//休眠10秒,测试令牌桶算法

	return true
}

// 释放一个连接
func (cl *ConnLimiter) ReleaseConn() {
	//从channel中读取一个
	c := <-cl.bucket

	log.Printf("释放连接%d", c)
}
