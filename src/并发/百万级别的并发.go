package main

import (
	"fmt"
	"time"
)

/*执行一次完整任务需要5秒*/
func DoTask(no int) {
	for i := 1; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(no)
	}
}
func main() {
	//主协程迅速开辟10条并发，耗时0.00000000000001秒
	for i := 1; i < 1000000; i++ {
		go DoTask(i)
	}
	//主协程打印完就死
	//主协程一死，子协程都陪着死
	time.Sleep(6 * time.Second)
	fmt.Println("main over")
}
