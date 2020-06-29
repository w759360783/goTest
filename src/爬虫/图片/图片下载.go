package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var (
	/*控制最大并发数的管道*/
	chSem = make(chan int, 5)
	/*主线程等待组*/
	wg sync.WaitGroup
	/*同步锁*/
	mu sync.Mutex
	/*下载文件存放的路径*/
	imgDir = `F:\golandWoke\src\爬虫\imgs\`
)

func main() {
	imgs := GetPageImgUrl("https://www.163.com")
	fmt.Println(imgs)
	for _, url := range imgs {
		fmt.Println(url)
		DownLoadImgAsync(url)
	}
}

/**
图片下载方法
 */
func DownLoadImg(img Image) {
	resp, err := http.Get(img.url)
	HandlerErr(err, "DownLoadImg:http.Get"+img.url)
	defer resp.Body.Close()
	imgBytes, err := ioutil.ReadAll(resp.Body)
	HandlerErr(err, "DownLoadImg:ioutil.ReadAll")
	var fileName string
	fileName = GetFileName(img, imgDir)
	err = ioutil.WriteFile(fileName, imgBytes, 0666)
	if err != nil {
		fmt.Println(fileName, "下载失败:", err)
	} else {
		fmt.Println(fileName, "下载成功")
	}
}

/**
异步下载
 */
func DownLoadImgAsync(img Image) {
	wg.Add(1)
	go func() {
		chSem <- 1
		DownLoadImg(img)
		wg.Done()
		<-chSem
	}()
	wg.Wait()
}
