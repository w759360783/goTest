package main

import (
	"fmt"
	"sync"
	"time"
)

/**
银行账户案例
·创建银行账户类
·存取款方法需要并发安全（不允许并发访问余额）
·查询余额和打印流水可以并发操作
·创建账户实力，并发执行存取款、查询余额、打印流水
 */

type Account struct {
	money int        // 余额
	mt    sync.Mutex // 账户的互斥锁
}

/**
存钱的操作
 */
func (a *Account) SaveMoney(n int) {
	a.mt.Lock()
	fmt.Println("SaveMoney---------start")
	<-time.After(3 * time.Second)
	a.money += n
	fmt.Println("SaveMoney---------end")
	a.mt.Unlock()
}

/**
取钱的操作
 */
func (a *Account) GetMoney(n int) {
	a.mt.Lock()
	fmt.Println("GetMoney---------start")
	<-time.After(3 * time.Second)
	a.money -= n
	fmt.Println("GetMoney---------end")
	a.mt.Unlock()
}

/**
查询当前余额
 */
func (a *Account) QueryMoney() {
	fmt.Println("QueryMoney-----------start")
	<-time.After(3 * time.Second)
	fmt.Println("当前余额为", a.money)
	fmt.Println("QueryMoney-----------end")

}

/**
打印流水
 */
func (a *Account) PrintHistory() {
	fmt.Println("PrintHistory-----------start")
	<-time.After(3 * time.Second)
	fmt.Println("打印流水")
	fmt.Println("PrintHistory-----------end")
}
func main() {
	var wg sync.WaitGroup
	account := &Account{1000,sync.Mutex{}}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			account.SaveMoney(100)
			wg.Done()
		}()
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			account.GetMoney(100)
			wg.Done()
		}()
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			account.QueryMoney()
			wg.Done()
		}()
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			account.PrintHistory()
			wg.Done()
		}()
	}
	wg.Wait()
}
