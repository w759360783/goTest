package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

/**
·发起百度搜索的GET请求："http://www.baidu.com/s?wd=肉丝"，打印回复的内容
·对https://httpbin.org/post发起post请求，携带表单数据，键值自拟，打印回复的内容
·表单数据类型application/x-www-form-urlencoded，表单读取API：strings.NewReader("rmb=0.5&hobby=接客和约汉"))
 */

func CHandlerError(err error, when string) {
	if nil != err {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
func main() {
	getDemo()
	//postDemo()
}

func getDemo() {
	resp, err := http.Get("http://www.baidu.com/s?wd=大后寿寿花")
	CHandlerError(err, "http.Get")
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	CHandlerError(err, "ReadAll")
	fmt.Println(string(bytes))
}

func postDemo() {
	resp, err := http.Post("https://httpbin.org/post?name=wu&age=21", "application/x-www-form-urlencoded",
		strings.NewReader("rmb=0.5&hobby=接客和约汉"))
	CHandlerError(err, "http.Post")
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	CHandlerError(err, "ReadAll")
	fmt.Println(string(bytes))
}
