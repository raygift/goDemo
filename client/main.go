package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	fmt.Println("ResolveTCPAddr done")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	fmt.Println("dialtcp done")
	checkError(err)
	wn, err := conn.Write([]byte("hello world\r\n\r\n"))
	fmt.Println("write done length: ", wn, err)
	checkError(err)
	content := make([]byte, 1024)
	time.After(1000)
	n, err := conn.Read(content)
	fmt.Println("Read done return: ", n, err)
	checkError(err)
	fmt.Println(string(content))
	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
