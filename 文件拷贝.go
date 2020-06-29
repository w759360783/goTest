package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*
√使用ioutil包做一个傻瓜式的拷贝
√使用io.Copy进行文件拷贝
√使用缓冲1K的缓冲区配合缓冲读写器进行文件拷贝
*/

// 使用ioutil包的文件拷贝
func main025() {
	bytes, err := ioutil.ReadFile("C:/Users/wuming/Pictures/Saved Pictures/there.png")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("C:/Users/wuming/Pictures/Saved Pictures/Copy there.png", bytes, 0666)
	if err != nil {
		panic(err)
	}
	fmt.Println("拷贝成功^_^")
}

//使用io.Copy进行文件拷贝
func main026() {
	src, err := os.OpenFile("C:/Users/wuming/Pictures/Saved Pictures/two.png", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer func() {
		src.Close()
		fmt.Println("文件src关闭成功")
	}()
	dst, err := os.OpenFile("C:/Users/wuming/Pictures/Saved Pictures/Copy two.png", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer func() {
		dst.Close()
		fmt.Println("文件dst关闭成功")
	}()
	written, err := io.Copy(dst, src)
	if err != nil {
		panic(err)
	}
	fmt.Println("复制成功,拷贝的字节数=", written)
}

//使用缓冲1K的缓冲区配合缓冲读写器进行文件拷贝
func main027() {
	//打开源文件
	src, err := os.OpenFile("C:/Users/wuming/Pictures/Saved Pictures/two.png", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	//打开目标文件
	dst, err := os.OpenFile("C:/Users/wuming/Pictures/Saved Pictures/Copy two.png", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	//延迟操作关闭文件资源
	defer func() {
		dst.Close()
		src.Close()
		fmt.Println("文件全部关闭成功")
	}()
	//创建文件读取缓冲器
	reader := bufio.NewReader(src)
	//创建文件写入缓冲器
	writer := bufio.NewWriter(dst)
	//创建临时缓冲区
	buf := make([]byte,1024)
	//循环读写
	for{
		//将文件读取到缓冲区
		_, err := reader.Read(buf)
		if err != nil {
			//表示源文件读取完毕退去循环
			if err == io.EOF{
				break
			}
			panic(err)
		}
		_, err = writer.Write(buf)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("拷贝成功")
}
