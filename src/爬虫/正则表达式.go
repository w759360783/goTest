package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

var (
	rePhone = `(1[3-9]\d)(\d{4})(\d{4})`
	// 759360783@qq.com
	//reEmail = `[\w\.]+@+[a-z]{2,5}\.(\.\w)?`
	reEmail = `[\w\.]+@\w+\.[a-z]{2,5}(\.\w)?`

	reUrl = `<a[\s\S]+?href="(http[\s\S]+?)"`
)

var (
	phoneUrl = `http://www.baishicha.com/hangzhou-yidong/`
	//emailUrl = `https://tieba.baidu.com/p/5366175841?red_tag=2053916567`
	emailUrl = `http://tieba.baidu.com/p/1634694280`

	url = "http://www.hao123.com"
)
/**
错误处理函数
 */
func ErrHandler(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}

/*手机号码爬取*/
func main1() {
	htmlStr := GetHtmlString(phoneUrl)
	compile := regexp.MustCompile(rePhone)
	/*普通的匹配*/
	//allString := compile.FindAllString(htmlStr, -1)
	//for _, value := range allString {
	//	fmt.Println(value)
	//}
	/*分组匹配*/
	submatch := compile.FindAllStringSubmatch(htmlStr, -1)
	for _, value := range submatch {
		fmt.Println(value)
	}
}

func GetHtmlString(url string) string {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	ErrHandler(err, "http.Get")
	bytes, err := ioutil.ReadAll(resp.Body)
	ErrHandler(err, "ReadAll")
	htmlStr := string(bytes)
	return htmlStr
}

/*邮箱爬取*/
func main2() {
	htmlString := GetHtmlString(emailUrl)
	compile := regexp.MustCompile(reEmail)
	allString := compile.FindAllString(htmlString, -1)
	for _, value := range allString {
		fmt.Println(value)
	}
}

/**
超链接爬取
 */
func main() {
	htmlString := GetHtmlString(url)
	compile := regexp.MustCompile(reUrl)
	allString := compile.FindAllStringSubmatch(htmlString, -1)
	for _, value := range allString {
		fmt.Println(value[1])
	}
}
