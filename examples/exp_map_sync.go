package main

import (
	"fmt"
	"sync"
)

func main() {
	// 声明 scene，类型为 sync.Map，注意，sync.Map 不能使用 make 创建。
	var scene sync.Map
	// 将键值对保存到sync.Map
	// 将一系列键值对保存到 sync.Map 中，sync.Map 将键和值以 interface{} 类型进行保存。
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	// 提供一个 sync.Map 的键给 scene.Load() 方法后将查询到键对应的值返回。
	fmt.Println(scene.Load("london")) // 100 true
	// 根据键删除对应的键值对
	// sync.Map 的 Delete 可以使用指定的键将对应的键值对删除。
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	// Range() 方法可以遍历 sync.Map，遍历需要提供一个匿名函数，参数为 k、v，类型为 interface{}，每次 Range() 在遍历一个元素时，都会调用这个匿名函数把结果返回。
	scene.Range(func(k, v interface{}) bool {
		// iterate: greece 97
		// iterate: egypt 200
		fmt.Println("iterate:", k, v)
		return true
	})
}

// sync.Map 没有提供获取 map 数量的方法，替代方法是在获取 sync.Map 时遍历自行计算数量，sync.Map 为了保证并发安全有一些性能损失，因此在非并发情况下，使用 map 相比使用 sync.Map 会有更好的性能。
