package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
服务器负载控制
·监听最大客户端连接数
·服务端协程：只要服务器过载，就通知控制协程，并进入阻塞等待
·控制协程：受到服务端预警，削减客户端数量，并通知服务端（预警已解除）
*/

type Server struct {
	maxConnection int        // 最大连接数
	chEarly       chan bool  // 预警通道
	sC            *sync.Cond // 预警锁
	connection    int        // 当前已经有的客户端连接数
}

// 初始化服务端
func NewServer(maxConnection int) *Server {
	server := new(Server)
	// 设置服务端最大连接数
	server.maxConnection = maxConnection
	// 初始化预警锁
	server.sC = sync.NewCond(&sync.Mutex{})
	// 初始化预警通道
	server.chEarly = make(chan bool)
	return server
}

// 服务端主协程
func (s *Server) Run() {
	fmt.Println("Run---------------------")
	for {
		s.sC.L.Lock()
		for s.maxConnection <= s.connection {
			// 发送最大荷载预警
			s.chEarly <- true
			fmt.Println("发送负载预警：各部门请注意，服务端已超过最大荷载！！！")
			// 等待荷载预警解除
			s.sC.Wait()
			fmt.Println("荷载预警解除---")
		}
		s.connection++
		// 模拟客户端连接耗时操作1s
		time.After(time.Second)
		fmt.Println("一个客户端加入了连接")
		s.sC.L.Unlock()
	}
}

// 控制协程当服务端主协程发送控制预警时随机削减服务端中的携程数量，并解除预警通知服务端主协程
func (s *Server) ServerControl() {
	fmt.Println("ServerControl---------------------")
	go func() {
		for {
			// 阻塞等待服务端主线程的荷载预警通知
			<-s.chEarly
			//收到服务端的预警通知
			fmt.Println("各部门：已经接收到服务端的预警通知----")
			s.sC.L.Lock()
			// 随机削减服务端连接池中的客户端数量,最大不超过服务端连接池的容量
			s.connection -= rand.Intn(s.maxConnection)
			// 模拟削减耗时操作 延迟3s
			<-time.After(3 * time.Second)
			fmt.Println("处理成功：警报解除当前连接池拥有", s.connection, "条连接", "空余连接", s.maxConnection-s.connection, "条")
			s.sC.Signal()
			s.sC.L.Unlock()
		}
	}()
}

func main() {
	server := NewServer(3)
	server.ServerControl()
	server.Run()
}
