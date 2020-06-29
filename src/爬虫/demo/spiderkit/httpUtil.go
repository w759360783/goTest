package spiderkit

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	DIAL_TIMEOUT = 10 * time.Second // 连接超时时间设置
	RW_TIMEOUT   = 10 * time.Second // 读写超时时间设置
)

var (
	client http.Client
)

/**
初始化方法
 */
func init() {
	client = http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				//设置连接请求超时时间
				conn, err := net.DialTimeout(netw, addr, DIAL_TIMEOUT)
				if err != nil {
					return nil, err
				}
				//设置连接的读写超时时间
				deadline := time.Now().Add(RW_TIMEOUT)
				err = conn.SetDeadline(deadline)
				if err != nil {
					return nil, err
				}
				return conn, nil
			},
		},
	}
}

/**
获取页面的html
 */

func GetHtmlString(url string) string {
	resp, err := client.Get(url)
	defer resp.Body.Close()
	HandlerErr(err, "GetHtmlString:client.Get")
	bytes, err := ioutil.ReadAll(resp.Body)
	return string(bytes)
}

/**
下载文件
 */
func DownloadFileWithClient(url, fileName string) {
	resp, err := client.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(fileName, "下载失败")
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fileName, "下载失败")
		return
	}
	err = ioutil.WriteFile(fileName, bytes, 0666)
	if err != nil {
		fmt.Println(fileName, "下载失败")
		return
	} else {
		fmt.Println(fileName, "下载成功")
	}
}
