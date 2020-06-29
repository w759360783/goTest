package main

import (
	"fmt"
	"time"
)

func main1() {
	nowTime := time.Now()
	fmt.Println(nowTime)
	//年月日
	////年
	//yeas := time.Now().Year()
	//fmt.Println(yeas)
	////月
	//month := time.Now().Month()
	//fmt.Println(month)
	////日
	//day := time.Now().Day()
	//fmt.Println(day)

	//今年的第几天
	//fmt.Printf("今天是今年的第%d天\n",time.Now().YearDay())
	////这个月的第几天
	//fmt.Printf("今天是这个月的第%d天\n",time.Now().Day())
	////这周的第几天
	//fmt.Printf("今天是这周的第%d天\n",time.Now().Weekday())

	//时分秒
	//fmt.Println(time.Now().Hour())
	//fmt.Println(time.Now().Minute())
	//fmt.Println(time.Now().Second())
	//fmt.Println(time.Now().Nanosecond())

	//创建时间
	date := time.Date(2019, time.February, 22, 15, 0, 0, 0, nowTime.Location())
	fmt.Println(date)
}

func main2() {
	now := time.Now()

	//一天之前
	duration, _ := time.ParseDuration("-24h0m0s")
	fmt.Println(now.Add(duration))
	//一周之前
	fmt.Println(now.Add(duration * 7))
	//一个月之前
	fmt.Println(now.Add(duration * 30))
}
