package spiderkit

import (
	"fmt"
	"strings"
)

type File struct {
	fileName string
	fileUrl  string
}

/*
获取页面上的全部信息
*/
func GetPageImginfos2Chan(url, imgDir string, chFile chan<- File) {
	// 获取当前url页面的所有html标签转换成的string
	htmlString := GetHtmlString(url)
	submatch := reImgTag.FindAllStringSubmatch(htmlString, -1)
	fmt.Println("已捕获图片", len(submatch), "张")

	for _, values := range submatch {
		chFile <- File{fileName: GetImgNameFromTag(values[0], values[1], imgDir), fileUrl: values[1]}
	}

}

/**
得到文件的文件名
 */
func GetImgNameFromTag(imgTag, imgUrl, imgDir string) string {
	// 从imgTag中获取图片的Alt属性中的内容
	submatch := reTagAlt.FindAllStringSubmatch(imgTag, -1)
	// 返回的文件名
	fileName := imgDir
	// 从图片的alt中获取的图片解释
	var altName string
	// 图片的后缀名,默认设置为jpg
	suName := ".jpg"
	// 图片下载链接中获取的文件名
	var imgName string
	if len(submatch) > 0 {
		altName = submatch[0][1]
	}
	/*根据图片的下载链接中的图片名获取图片的后缀名*/
	// 1.从图片的下载连接中获取图片名
	imgNameSubmatch := reImgName.FindAllStringSubmatch(imgUrl, -1)
	if len(imgNameSubmatch) > 0 {
		imgName = imgNameSubmatch[0][1]
		suName = imgName[strings.LastIndex(suName, "."):]
	}
	if altName != "" || imgName != "" {
		//优先的命名规范为altName+"_"+imgName
		if imgName == "" {
			fileName += altName + suName
		} else {
			fileName += altName + "_" + imgName
		}
	} else {
		// 次选用时间戳+"_"+随机数命名
		fileName += GetRandomFileName() + suName
	}
	return fileName
}
