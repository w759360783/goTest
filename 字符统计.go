package main

import (
	"fmt"
	"io/ioutil"
	"unicode/utf8"
)

func main (){
	bytes, err := ioutil.ReadFile("C:/Users/wuming/Desktop/读写文件测试1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
	fmt.Printf("读取的文件中的字符数为%d",utf8.RuneCountInString(string(bytes)))
}
