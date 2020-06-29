package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
·开辟三条协程：
·A每秒生成一个三位随机数
·B输出该值的奇偶性
·C输出奇数和偶数的总个数；
·当生成一个水仙花数时，程序结束；
*/
func main() {
	// 生成的随机数传输管道
	chNumber := make(chan int, 3)
	// 随机数的性质和值的传输管道
	chValueParity := make(chan [2]interface{}, 3)
	//控制程序退出的管道
	chExit := make(chan int)
	/**
	不停的输出随机数的协程
	 */
	go func() {
		for {
			random := GetRandom()
			ticker := time.NewTicker(time.Microsecond)
			fmt.Println("生成了", random)
			chNumber <- random
			<-ticker.C
		}
	}()
	/*
	判断输出的随机数的奇偶性
	*/
	go func() {
		for i := range chNumber {
			even := IsEven(i)
			if even {
				fmt.Println(i, "是偶数")
			} else {
				fmt.Println(i, "是奇数")
			}
			chValueParity <- [2]interface{}{i, even}
		}
	}()

	/**
	输出奇数和偶数的总个数，并且在碰到生成的数为水仙花数的时候控制程序的退出
	 */
	go func() {
		NumberEvenOdd(chValueParity, chExit)
	}()
	<-chExit
}

/**
生成一个随机的3位随机数
 */
func GetRandom() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret := 100 + r.Intn(999-100+1)
	return ret
}

/*
输出传入的值的奇偶性
*/
func IsEven(n int) bool {
	return n%2 == 0
}

/**
输出奇偶数的个数
 */
func NumberEvenOdd(chValueParity chan [2]interface{}, chExit chan int) {
	var evenCount, OddCount int
	for i := range chValueParity {
		if i[1].(bool) {
			evenCount ++
		} else {
			OddCount ++
		}
		fmt.Println("奇数的个数=", OddCount, "偶数的个数=", evenCount)
		if IsNarcissistic(i[0].(int)) {
			fmt.Println(i[0], "是一个水仙花数")
			chExit <- -1
		}
	}
}

/**
判断一个数是否为水仙花数
 */
func IsNarcissistic(n int) bool {
	a := n / 100
	b := n % 100 / 10
	c := n % 10
	return a*a*a+b*b*b+c*c*c == n
}
