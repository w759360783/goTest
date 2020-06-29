package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			runtime.Gosched()
		}
		go func(number int) {
			for j := 0; j < 10; j++ {
				fmt.Printf("协程%d：%d\n", number, j)
			}
		}(i)
	}
	time.Sleep(3 * time.Second)
}
