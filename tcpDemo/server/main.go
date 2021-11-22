package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "192.168.92.3:50051")
	if err != nil {
		fmt.Println("listen error: ", err)
		return
	} else {

		fmt.Println("listen done")
		fmt.Println(listen.Addr().String())
	}

	for {
		conn, err := listen.Accept()
		fmt.Println("accept done")
		if err != nil {
			fmt.Println("accept error: ", err)
			break
		}

		// start a new goroutine to handle the new connection
		go HandleConn(conn)
	}
}
func HandleConn(conn net.Conn) {
	fmt.Println("HandleConn start")
	defer conn.Close()
	packet := make([]byte, 1024)
	// for {
	// 如果没有可读数据，也就是读 buffer 为空，则阻塞
	n, err := conn.Read(packet)
	fmt.Println("Read return: ", n, err)
	// 同理，不可写则阻塞
	n, err = conn.Write(packet)
	fmt.Println("Write return: ", n, err)

	// }
}
