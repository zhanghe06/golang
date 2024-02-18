package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

var tcpCHost = flag.String("host", "localhost", "host")
var tcpCPort = flag.String("port", "9999", "port")

type tcpCMsg struct {
	Data string `json:"data"`
	Type int    `json:"type"`
}
type tcpCResp struct {
	Data   string `json:"data"`
	Status int    `json:"status"`
}

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *tcpCHost+":"+*tcpCPort)
	if err != nil {
		fmt.Println("connect error", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("connecting to " + *tcpCHost + ":" + *tcpCPort)
	var wg sync.WaitGroup
	wg.Add(2)
	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)
	wg.Wait()
}

func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 10; i > 0; i-- {
		d := "hello" + strconv.Itoa(i)
		msg := tcpCMsg{
			Data: d,
			Type: 1,
		}
		b, _ := json.Marshal(msg)
		writer := bufio.NewWriter(conn)
		writer.Write(b)
		writer.Write([]byte("\n"))
		writer.Flush()
	}
	fmt.Println("write done")
}

func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(conn)
	for i := 1; i <= 10; i++ {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("read error", err)
			return
		}
		var resp tcpCResp
		json.Unmarshal(line, &resp)
		fmt.Println("status", resp.Status, " content:", resp.Data)
	}
	fmt.Println("read done")
}
