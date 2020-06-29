package main

import "fmt"

func main() {
	fmt.Println("《望庐山瀑布》")
	fmt.Println("作者：李白")
	fmt.Println("日照香炉生紫烟，")
	fmt.Println("遥看瀑布忘前川。")
	fmt.Println("飞流直下三千尺，")

	goto OVER
	fmt.Println("疑是银河落九天。")

OVER:
	fmt.Println("俺是乱入的")
}
