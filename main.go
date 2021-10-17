package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"time"
)

const (
	localhost   = "192.168.92.3"
	defaultPort = 20000
	port2       = 20001
)

var defaultAddr = fmt.Sprintf("%s:%d", localhost, defaultPort)
var count int = 0

// TCP
func serverProcess(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		_, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("server: read from client failed, err:", err)
			break
		}
		// fmt.Println("server: recv - ", string(buf[:n]))
		reply := "hello client"
		conn.Write([]byte(reply)) // 发送数据
	}
}

func startServer(addr string) {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("server: listen failed, err:", err)
		return
	}
	fmt.Println("server: start listening")
	go func(listen net.Listener) {
		for {
			conn, err := listen.Accept() // 监听连接
			if err != nil {
				fmt.Println("server: accept failed, err:", err)
				continue
			}
			fmt.Println("server: accepted new client")
			go serverProcess(conn) // 启动一个goroutine处理连接
		}
	}(listen)
}

func startClient(addr string) {
	var conn net.Conn
	conn, err := net.Dial("tcp", addr)
	if err != nil { // 建立连接
		fmt.Println("err :", err)
		return
	}
	fmt.Println("client: connected to server")
	go func(conn net.Conn) {
		defer conn.Close() // 关闭连接
		for {
			message := "hello server"
			_, err := conn.Write([]byte(message)) // 发送数据至server
			if err != nil {
				return
			}

			buf := [128]byte{}
			_, err = conn.Read(buf[:]) // 读取server返回数据
			if err != nil {
				fmt.Println("client: recv failed, err:", err)
				return
			}
			// fmt.Println("client: recv - " + string(buf[:n]))
			count++
		}
	}(conn)
}

func localTCP() {
	startServer(defaultAddr)
	startClient(defaultAddr)
	time.Sleep(time.Duration(200) * time.Millisecond)
	fmt.Printf("total: %d\n", count)
}

func main() {
	isLocal := flag.Bool("l", false, "set if test local tcp")
	isServer := flag.Bool("s", false, "set if act as server")

	addr := flag.String("a", "", "server address")
	flag.Parse()

	if *isLocal {
		localTCP()
	} else {
		if *isServer {
			startServer(defaultAddr)
			for {
				time.Sleep(time.Duration(20) * time.Second)
			}

		} else {
			startClient(fmt.Sprintf("%s:%d", *addr, defaultPort))
			time.Sleep(time.Duration(2000) * time.Millisecond)
			fmt.Printf("total: %d\n", count)
		}
	}
}
