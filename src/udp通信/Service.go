package main

import (
	"fmt"
	"net"
	"os"
)

func ServiceHandlerError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
	ServiceHandlerError(err, "ResolveTCPAddr")
	conn, err := net.ListenUDP("udp", addr)
	ServiceHandlerError(err, "ListenUDP")
	buffer := make([]byte, 1024)
	for {
		n, udpAddr, err := conn.ReadFromUDP(buffer)
		ServiceHandlerError(err, "Read")
		ClientMsg := string(buffer[:n])
		fmt.Printf("来自%v的消息:%s", udpAddr, ClientMsg)
		_, err = conn.WriteToUDP([]byte("朕已阅"),udpAddr)
		ServiceHandlerError(err, "Write")
	}
}
