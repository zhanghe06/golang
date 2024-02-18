package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var tcpSHost = flag.String("host", "", "host")
var tcpSPort = flag.String("port", "9999", "port")

type tcpSMsg struct {
	Data string `json:"data"`
	Type int    `json:"type"`
}

type tcpSResp struct {
	Data   string `json:"data"`
	Status int    `json:"status"`
}

func handleRequest(conn net.Conn) {
	ip := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnect:" + ip)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	for {
		b, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		var msg tcpSMsg
		json.Unmarshal(b, &msg)
		fmt.Println("GET==>data ", msg.Data, " type:", msg.Type)
		resp := tcpSResp{
			Data:   time.Now().String(),
			Status: 200,
		}
		r, _ := json.Marshal(resp)
		writer.Write(r)
		writer.Write([]byte("\n"))
		writer.Flush()
	}
	fmt.Println("done!")
}

func main() {
	flag.Parse()
	var l net.Listener
	var err error
	l, err = net.Listen("tcp", *tcpSHost+":"+*tcpSPort)
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("listening on " + *tcpSHost + ":" + *tcpSPort)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			os.Exit(1)
		}
		fmt.Printf("message %s->%s\n", conn.RemoteAddr(), conn.LocalAddr())
		go handleRequest(conn)
	}
}
