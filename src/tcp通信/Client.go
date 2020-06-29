package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

/**
·拨号连接服务端主机的8888端口，建立连接
·循环从标准输入（命令行）读取一行用户输入，向服务端发送
·接收并打印服务端消息，如果消息是“bye”，就退出程序
 */

//错误处理的封装
func ClientHandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	ClientHandleError(err, "Dial")
	buffer := make([]byte, 1024)
	read := bufio.NewReader(os.Stdin)
	for {
		line, _, _ := read.ReadLine()
		_, err := conn.Write(line)
		ClientHandleError(err, "Write")

		n, err := conn.Read(buffer)
		ClientHandleError(err, "Read")
		ServiceMsg := string(buffer[:n])
		fmt.Println("服务端:",ServiceMsg)
		if ServiceMsg == "bye" {
			break
		}
	}
	fmt.Println("GG")
}
