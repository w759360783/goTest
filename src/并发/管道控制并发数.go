package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func getSqrt(name string, n int, semChan chan string, exit chan<- int) {
	semChan <- name
	sqrt := math.Sqrt(float64(n))
	fmt.Printf("%s:%d的平方根为%v\n", name, n, sqrt)
	<-semChan
	exit <- 1
}
func main() {
	semChan := make(chan string, 5)
	exit := make(chan int)
	for i := 0; i < 100; i++ {
		go getSqrt("协程"+strconv.Itoa(i), i, semChan, exit)
	}

	for {
		time.Sleep(time.Second * 1)
	}

}
