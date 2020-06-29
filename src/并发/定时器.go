package main

import (
	"fmt"
	"os"
	"time"
)

func main015() {
	timer := time.NewTimer(time.Second * 5)
	fmt.Println("计时开始当前时间为", time.Now())
	<-timer.C
	fmt.Println("时间已到当前时间为", time.Now())
}

/*定时器的终止*/
func main016() {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("计时开始", time.Now())
	go func() {
		time.Sleep(3 * time.Second)
		// 停止计时器（永远不会向Timer.C中写入数据）
		ok := timer.Stop()
		if ok {
			fmt.Println("炸弹已拆除")
			os.Exit(0)
		}
	}()
	// 阻塞5s
	endTime := <-timer.C
	fmt.Println("炸弹爆炸于", endTime)
}

/*定时器的重置*/
func main() {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("计时开始", time.Now())
	time.Sleep(2 * time.Second)
	timer.Reset(2 * time.Second)

	// 阻塞5秒
	endTime := <-timer.C
	fmt.Println("炸弹爆炸于", endTime)
}
