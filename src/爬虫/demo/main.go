package demo

import (
	"spiderkit"
	"strconv"
	"sync"
)

var (
	// 控制最大下载并发数的通道
	chSem = make(chan int, 10)
	// 控制获取页面资源的传输通道
	chchImgMaps = make(chan spiderkit.File, 1000)

	// 等待图片解析
	wgImginfo sync.WaitGroup
	// 等待图片下载
	wgDownload sync.WaitGroup

	// 图片下载地址
	imgDir = `F:\golandWoke\src\爬虫\demo\img\`
)

func main() {
	// 被爬取的网站的基类地址
	baseUrl := "https://www.duotoo.com/zt/rbmn/index"
	for i := 1; i < 10; i++ {
		/*获取当前需要请求的地址*/
		url := baseUrl
		if i != 1 {
			url += "_" + strconv.Itoa(i)
		}
		url += ".html"

		// 解析地址获取待下载的图片资源文件
		go func(srcUrl string) {
			spiderkit.GetPageImginfos2Chan(srcUrl, )
		}(url)
	}
}
