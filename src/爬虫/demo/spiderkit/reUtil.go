package spiderkit

import "regexp"

/**
图片爬取的正则表达式的工具类
 */

var (
	/*从html标签中抓取图片元素*/
	reImgTagStr = `<img.+?src="(http.+?)".*?>`
	/*抓取图片中的alt标签中的内容*/
	reImgAltStr = `<img.+?alt="(.+?)"`
	/*从已经抓取的img标签中提取alt标签中的内容*/
	reTagAltStr = `alt="(.+?)"`
	/*img标签中的src元素*/
	reTagSrcStr = `src="(http.+?)"`
	/*抓取图片下载连接中的文件名*/
	reImgNameStr = `/(\w+\.((jpg)|(jpeg)|(png)|(gif)|(bmp)|(webp)|(swf)|(ico)))`
	/*预编译正则对象*/
	reImgTag, reImgAlt, reTagAlt, reTagSrc, reImgName *regexp.Regexp
)

func init() {
	reImgTag = regexp.MustCompile(reImgTagStr)
	reImgAlt = regexp.MustCompile(reImgAltStr)
	reTagAlt = regexp.MustCompile(reTagAltStr)
	reTagSrc = regexp.MustCompile(reTagSrcStr)
	reImgName = regexp.MustCompile(reImgNameStr)
}
