package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

const (
	col = 10000
	row = 10000
)

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil { //监控cpu
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// 主逻辑区，进行一些简单的代码运算
	x := [row][col]int{}
	s := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			x[i][j] = s.Intn(100000)
		}
	}

	for i := 0; i < row; i++ {
		tmp := 0
		for j := 0; j < col; j++ {
			tmp += x[i][j]
		}
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC()                                      // GC，获取最新的数据信息
		if err := pprof.WriteHeapProfile(f); err != nil { // 写入内存信息
			log.Fatal("could not write memory profile: ", err)
		}
		_ = f.Close()
	}
}

// go run exp_pprof.go -cpuprofile cpu.prof -memprofile mem.prof
// go tool pprof
// go tool pprof cpu.prof
// go tool pprof mem.prof

// (pprof) web

// MacOS安装Graphviz
// brew install graphviz
