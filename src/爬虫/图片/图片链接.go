package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	/*
	<img src="http://cms-bucket.ws.126.net/2019/05/07/f18d9fb14f5e4ec3825edbe18840488a.jpeg?imageView&amp;thumbnail=200y125">
	 */
	//reImg = `<img[\s\S]+?src="(http[\s\S]+?)"`
	//reImg = `<img.*?src="(http.+?)".*?(alt="(.*?)")?/>`
	// 获取所有的Img的标签
	reImg = `<img.+?src="(http.+?)".*?>`
	// 从Img标签中提取出alt中的内容
	reExtractImg = `alt="(.+?)"`
	// 从图片的下载链接中提取出文件名
	reDownloadName = `/(\w+\.((jpg)|(jpeg)|(png)|(gif)|(bwp)))`
)

/**
图片结构体
 */
type Image struct {
	name string
	url  string
}

/*
从图片的下载链接中获取图片的文件名
*/
func getDownloadUrlName(downloadUrl string) string {
	compile := regexp.MustCompile(reDownloadName)
	submatch := compile.FindAllStringSubmatch(downloadUrl, -1)
	fmt.Println(submatch)
	if len(submatch) > 0 {
		return submatch[0][1]
	} else {
		return ""
	}
}

/**
从图片标签当中获取到Alt中的内容
 */
func getAlt(imgTag string) string {
	compile := regexp.MustCompile(reExtractImg)
	allString := compile.FindAllStringSubmatch(imgTag, -1)
	//fmt.Println(allString)
	//fmt.Println(len(allString))
	if len(allString) > 0 {
		return allString[0][1]
	}
	return ""
}

/**
获得页面上的全部图片链接
"https://www.163.com"
 */
func GetPageImgUrl(url string) []Image {
	s := getHtmlString(url)
	// 打印整体的html的内容
	compile := regexp.MustCompile(reImg)
	submatch := compile.FindAllStringSubmatch(s, -1)
	fmt.Println(len(submatch))
	images := make([]Image, 0)
	for _, value := range submatch {
		images = append(images, Image{getAlt(value[0]), value[1]})
	}

	for _, value := range images {
		fmt.Println(value)
	}
	return images
}

/**
下载到的图片命名规则方法
alt>图片路径中文件名>当前时间+随机数的文件名
 */
func GetFileName(image Image, dir string) string {
	var fileName string
	fileName += dir
	su := ".jpg"
	name := getDownloadUrlName(image.url)
	// 假如该图片的下载地址中有文件的后缀名则提取出来沿用该图片的原始文件名
	if name != "" {
		su = name[strings.LastIndex(name, "."):]
	}
	if name := getDownloadUrlName(image.url); (name != "") || (image.name != "") {
		image.name = strings.Replace(image.name, ":", "_", -1)
		if name == "" {
			fileName += image.name + su
		} else {
			fileName += image.name + "_" + name
		}

	} else {
		fileName += getRandomName() + su
	}
	return fileName
}

/**
错误处理函数
 */
func HandlerErr(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}

/**
请求链接地址将访问的Html的标签转换成字符串
 */
func getHtmlString(url string) string {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	HandlerErr(err, "http.Get")
	reader := simplifiedchinese.GB18030.NewDecoder().Reader(resp.Body)
	//bytes, err := ioutil.ReadAll(resp.Body)
	bytes, err := ioutil.ReadAll(reader)
	HandlerErr(err, "ioutil.ReadAll")
	return string(bytes)
}

/**
生成当前时间加随机数的命名
 */
func getRandomName() string {
	s := strconv.Itoa(int(time.Now().UnixNano()))
	s1 := strconv.Itoa(getRandomNumb(1000, 10000))
	return s + "_" + s1
}

/**
生成随机数
 */
func getRandomNumb(start, end int) int {
	mu.Lock()
	<-time.After(time.Nanosecond)
	i := rand.New(rand.NewSource(time.Now().UnixNano()))
	mu.Unlock()
	return start + i.Intn(end-start)

}
