package main

import (
	"fmt"
	"sync"
	"time"
)

/*
·创建40条协程，30条模拟读取数据库，10条模拟写入数据库
·最大并发数控制在5
·数据库允许一写多读
·使用定时器和秒表模拟IO耗时
·主协程必须恰好结束在所有协程完成工作的时候
*/

func main() {
	// 等待组
	var wg sync.WaitGroup
	// 读写锁
	var rwm sync.RWMutex
	//控制最大并发数的管道
	chSem := make(chan int, 5)

	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			wg.Add(1)
			go Read(&wg, chSem, &rwm)
		}
		wg.Add(1)
		go Write(&wg, chSem, &rwm)
	}
	// 等待所有的协程完成再退出程序
	wg.Wait()
}

func Read(wg *sync.WaitGroup, chSem chan int, rwm *sync.RWMutex) {
	rwm.RLock()
	fmt.Println("Read")
	chSem <- 123
	/*模拟IO耗时操作*/
	timer := time.NewTimer(3 * time.Second)
	<-timer.C
	wg.Done()
	<-chSem
	rwm.RUnlock()
}

func Write(wg *sync.WaitGroup, chSem chan int, rwm *sync.RWMutex) {
	rwm.Lock()
	fmt.Println("Write")
	chSem <- 123
	/*模拟IO耗时操作*/
	timer := time.NewTimer(3 * time.Second)
	<-timer.C
	wg.Done()
	<-chSem
	rwm.Unlock()
}
