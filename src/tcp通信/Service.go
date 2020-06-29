package main

import (
	"fmt"
	"net"
	"os"
)

/*
·服务端在本机的8888端口建立TCP监听
·为接入的每一个客户端开辟一条独立的协程
·循环接收客户端消息，不管客户端说什么，都自动回复“已阅xxx”
·如果客户端说的是“im off”，则回复“bye”并断开其连接
----------
·拨号连接服务端主机的8888端口，建立连接
·循环从标准输入（命令行）读取一行用户输入，向服务端发送
·接收并打印服务端消息，如果消息是“bye”，就退出程序
*/

//错误的封装
func CHandlerError(err error, when string){
	if err != nil {
		fmt.Println(err,when)
		os.Exit(1)
	}
}

func main (){
	//设置tcp监听器
	Listener ,err := net.Listen("tcp","127.0.0.1:8888")
	fmt.Println("Listener......")
	//错误处理
	CHandlerError(err,"Listen")
	//无限循环等待客户端连接
	for{
		//阻塞等待客户端连接
		conn, err := Listener.Accept()
		//错误处理
		CHandlerError(err,"Accept")
		//给连接上的客户端开启新goroutine
		go ReadAndWrite(conn)
	}
}

//与客户端的交互操作
func ReadAndWrite(conn net.Conn){
	buffer := make([]byte,1024)
	for{
		n, err := conn.Read(buffer)
		CHandlerError(err,"Read")
		clientMsg := string(buffer[:n])
		fmt.Printf("收到%v的消息:%s\n",conn.RemoteAddr(),clientMsg)

		if clientMsg != "im off"{
			_, err := conn.Write([]byte("已阅:" + clientMsg))
			CHandlerError(err,"Write")
		}else {
			_, err := conn.Write([]byte("bye"))
			CHandlerError(err,"Write")
			break
		}
	}
	//断开客户端连接
	err := conn.Close()
	CHandlerError(err,"Close")
	fmt.Printf("客户端%s已经断开连接\n", conn.RemoteAddr())
}
