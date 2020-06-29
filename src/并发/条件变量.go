package main

import (
	"fmt"
	"sync"
	"time"
)

/*
投资股票的条件变量示例
*/
func main() {
	// 股票的涨跌标志
	sharesFlag := false
	var cond = sync.NewCond(&sync.Mutex{})
	// 模仿市场的股票涨跌
	go func() {
		for {
			// 市场循环改变股票的涨跌，并且在改变的时候锁定涨跌的标志，防止并发错误,每两秒钟改变一次
			ticker := time.NewTicker(2 * time.Second)
			cond.L.Lock()
			sharesFlag = !sharesFlag
			// 向市场广播当前股票的涨跌状态已经改变
			cond.Broadcast()
			cond.L.Unlock()
			// 每两秒钟循环一次
			<-ticker.C
		}
	}()
	// 主线程模拟股民的对股票的操作
	for {
		cond.L.Lock()
		for !sharesFlag {
			// 当前不是牛市
			fmt.Println("当前不是牛市，请止损抛售")
			cond.Wait()
		}
		fmt.Println("股票涨了呀，超级大牛市，买买买")
		cond.L.Unlock()
	}
}
