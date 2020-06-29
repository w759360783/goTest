package spiderkit

import (
	"strconv"
	"time"
)

/**
生成时间戳_随机数名
 */

func GetRandomFileName() string {
	// 当前时间戳
	timeStamp := strconv.Itoa(int(time.Now().UnixNano()))
	// 随机数名
	randomName := GetRandomInt(1000, 10000)

	return string(timeStamp) + "_" + string(randomName)
}
