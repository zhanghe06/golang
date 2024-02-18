package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func loopPut(client *clientv3.Client) {
	ctx := context.TODO()
	c := 0
	for {
		c++
		time.Sleep(time.Second)
		res, _ := client.Put(ctx, fmt.Sprintf("k/%d", c), "v")
		fmt.Println(res)
	}
}

func loopGet(client *clientv3.Client) {
	ctx := context.TODO()
	res, _ := client.Get(ctx, "k", clientv3.WithPrefix())
	fmt.Println(res)
	if res.Count == 0 {
		fmt.Println("result is empty")
		return
	}
	for i, v := range res.Kvs {
		fmt.Printf("i: %d, k: %s, v: %s\n", i, v.Key, v.Value)
	}
}

func kvPut(client *clientv3.Client) {
	// 实例化一个用于操作ETCD的KV
	kv := clientv3.NewKV(client)
	putResp, err := kv.Put(context.TODO(), "/school/class/students", "tom", clientv3.WithPrevKV())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(putResp.Header.Revision)
	if putResp.PrevKv != nil {
		fmt.Printf(" prev Value: %s\n CreateRevision: %d\n ModRevision: %d\n Version: %d\n",
			string(putResp.PrevKv.Value), putResp.PrevKv.CreateRevision, putResp.PrevKv.ModRevision, putResp.PrevKv.Version)
	}
}

func kvGet(client *clientv3.Client) {
	// 实例化一个用于操作ETCD的KV
	kv := clientv3.NewKV(client)
	getResp, err := kv.Get(context.TODO(), "/school/class/students")
	if err != nil {
		fmt.Println(err)
		return
	}
	if getResp.Count == 0 {
		fmt.Println("result is empty")
		return
	}
	// 输出本次的Revision
	fmt.Printf("K: %s \nV: %s", getResp.Kvs[0].Key, getResp.Kvs[0].Value)
}

func main() {
	// 客户端配置
	config := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	// 建立连接
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	//loopPut(client)
	loopGet(client)

	//for {
	//	time.Sleep(time.Second)
	//	kvPut(client)
	//}

	// 写入
	//kvPut(client)

	// 读取
	//kvGet(client)
}

// kvPut
// 1177666
//  prev Value: lily
//  CreateRevision: 1177661
//  ModRevision: 1177665
//  Version: 5

// kvGet
// K: /school/class/students
// V: tom
