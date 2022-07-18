package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		os.Exit(1)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			return
		}
		go ChatWithConn(conn)
	}
}

func ChatWithConn(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("读取流出错")
			break
		}

		msg := string(buffer[0:n])
		fmt.Println("读取到消息：", msg)
		if msg == "off" {
			conn.Write([]byte("再见"))
			break
		}

		conn.Write([]byte("已读"))
		conn.Close()
		fmt.Printf("客户端断开连接: %s", conn.RemoteAddr())
	}
}
