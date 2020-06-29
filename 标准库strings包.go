package main

import (
	"fmt"
	"strings"
)

//检索字符串
func main03() {
	//fmt.Println(strings.Contains("hello","el")) // true
	//fmt.Println(strings.Contains("hello","aa")) // false
	//fmt.Println(strings.ContainsAny("hello","aaasdfe")) // ture
	//fmt.Println(strings.ContainsAny("hello","piu")) // fase

	//fmt.Printf("%d\n","我")
	//fmt.Println(strings.ContainsRune("爱上一只猪",'猪'))

}

//格式化
func main04() {
	fmt.Println(strings.ToLower("Hello"))
	fmt.Println(strings.ToUpper("Hello"))
	fmt.Println(strings.Title("hello"))
}

//比较大小
func main05() {
	fmt.Println(strings.Compare("wuming", "ztt"))
}

//裁剪
func main06() {
	////去掉首尾的空格
	//fmt.Println(strings.TrimSpace("   love a pig   "))
	////去掉前缀
	//fmt.Println(strings.TrimPrefix("wlove a pig","w"))
	////去掉后缀
	//fmt.Println(strings.TrimSuffix("love a pigm","m"))

	////去掉首尾的字符串
	//fmt.Println(strings.Trim("w爱上一只猪m","wm"))
	////去掉左边开始的字符串
	//fmt.Println(strings.TrimLeft("w爱上一只猪m","wm"))
	////去掉右边开始的字符串
	//fmt.Println(strings.TrimRight("w爱上一只猪m","wm"))
}

//炸碎
func main07() {
	//strs := strings.Split("love_a_pig", "_")
	//fmt.Println(strs, len(strs))
	//
	////控制切割成几串当n的为-1是和上面效果一致
	//strs = strings.SplitN("love_a_pig", "_", 2)
	//fmt.Println(strs, len(strs))
	//for _, v := range strs {
	//	fmt.Println(v)
	//}
	//
	//strs = strings.SplitAfter("love_a_pig", "_")
	//fmt.Println(strs, len(strs))
}

//拼接
func main08(){
	fmt.Println(strings.Join([]string{"wm","ztt"}," ❤ "))
}
