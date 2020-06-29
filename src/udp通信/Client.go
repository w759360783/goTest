package main

import (
	"fmt"
	"net"
	"os"
)

func ClientHandlerError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8888")
	ClientHandlerError(err, "Dial")
	buffer := make([]byte, 1024)
	n, err := conn.Write([]byte("啊哈哈哈，我是向服务端传输的额数据呀"))
	ClientHandlerError(err, "Write")
	fmt.Printf("向服务端传输%d字节的数据", n)
	n, err = conn.Read(buffer)
	ClientHandlerError(err, "Read")
	fmt.Println("服务端：", string(buffer[:n]))
}
