package main

import (
	"fmt"
	"sync"
	"time"
)

/*
城管预警
·监听城管大队
·烧烤摊集群：
	监听城管大队，
	只要出动就发消息通知工会主席并进入阻塞等待至被唤醒
	否则就提供露天烧烤
·工会主席：摆平城管大队，并广播通所有烧烤摊主
*/

var cityManagement = false // 标志城管大队是否出动的标志

var cond *sync.Cond // 条件变量

var chMsg chan bool // 通知通道

func main() {
	cond = sync.NewCond(&sync.Mutex{})
	chMsg = make(chan bool)
	go unionChairman()
	go cityManagements()
	Barbecue()
}

/**
烧烤摊贩
 */
func Barbecue() {
	fmt.Println("barbecue-------------start")
	for {
		cond.L.Lock()
		if cityManagement {
			chMsg <- true
			fmt.Println("城管已出动，我们做不了生意了，请主席给我们做主呀")
			cond.Wait()
		}
		fmt.Println("啦啦啦~我是快乐的小摊贩~啦啦啦我在烧烤")
		<-time.After(1 * time.Second)
		cond.L.Unlock()
	}
}

/**
公会主席
 */
func unionChairman() {
	fmt.Println("unionChairman-------")
	fmt.Println("我是很牛逼的工会主席，我正在度假~~~~")
	for {
		// 堵塞等待被窝庇护的摊贩的求助消息
		<-chMsg
		// 当前被我庇护的摊贩向我求助弄走城管
		cond.L.Lock()
		// 驱逐城管大队,模拟3s钟的耗时操作
		<-time.After(3 * time.Second)
		cityManagement = false
		// 广播通知商贩城管已经被弄走 可以出来摆摊了
		cond.Broadcast()
		fmt.Println("各摊贩请注意，警报已解除！")
		cond.L.Unlock()
	}
}

/**
城管
 */
func cityManagements() {
	fmt.Println("cityManagements------start")
	ticker := time.NewTicker(3 * time.Second)
	for {
		if !cityManagement {
			cond.L.Lock()
			// 每3秒钟城管出来巡逻一次
			fmt.Println("小摊贩出来受死，城管巡逻")
			cityManagement = true
			cond.L.Unlock()
		}
		<-ticker.C
	}
}
