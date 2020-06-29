package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	answer := myRand.Intn(1000)

	// 接受用户的输入
	for {
		fmt.Println("请输入您的猜想：")
		var guess string
		fmt.Scan(&guess)

		// 将用户输入的数转换成整数
		guessNum, err := strconv.Atoi(guess)
		if err != nil {
			fmt.Println(err)
		}
		if guessNum > answer {
			fmt.Println("大了")
		} else if guessNum < answer {
			fmt.Println("小了")
		} else {
			fmt.Println("恭喜您答对了")
			return
		}
	}
}
