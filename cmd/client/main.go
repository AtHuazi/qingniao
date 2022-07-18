package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	buffer := make([]byte, 1024)
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("读取用户输入错误")
			break
		}
		conn.Write(line)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("读取服务端返回流出错")
			break
		}
		msgReply := string(buffer[0:n])
		fmt.Println("服务端返回消息：", msgReply)
		if msgReply == "再见" {
			break
		}
	}
}
