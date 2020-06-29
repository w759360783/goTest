package main

import (
	"fmt"
	"time"
)

func main() {
	liuKe := ke("刘玄德", 0)
	zhuKe := ke("诸葛孔明", 2)
	pangKe := ke("庞统", 3)

	go func() {
		for i := 0; i < 7; i++ {
			result := liuKe()
			fmt.Println(result)
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for i := 0; i < 7; i++ {
			result := zhuKe()
			fmt.Println(result)
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for i := 0; i < 7; i++ {
			result := pangKe()
			fmt.Println(result)
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(8 * time.Second)
}

var wujiang = []string{"关羽", "张飞", "赵云", "马超", "黄忠"}

func ke(zhugong string, start int) func() string {
	index := start
	result := func() string {
		wuName := wujiang[index]
		index++
		if index > 4 {
			index = 0
		}
		return zhugong + "麾下：" + wuName + "出战"
	}
	return result
}
