package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//缓冲文件读取
func main20() {
	//0666表示所有人都对文件有读写的权限
	//4读取权限2写入权限1执行权限
	file, err := os.OpenFile("C:/Users/wuming/Desktop/读写文件测试.txt", os.O_RDONLY, 0666)
	//延迟函数关闭资源
	defer func() {
		fmt.Println("文件已关闭!")
		file.Close()
	}()
	//判断文件打开操作执行的结果
	if err == nil {
		fmt.Println("文件打开成功")
	} else {
		fmt.Println("文件打开失败err:", err)
		return
	}
	//创建文件的缓冲读取器
	read := bufio.NewReader(file)
	for {
		//表示每次读取到换行符就返回效果相当于逐行读取
		str, err := read.ReadString('\n')
		//判断读出是否成功，如果error为nil表示没有错误，则表示读取文件成功
		if err != nil {
			//错误当中有一个特殊的错误当错误为io.EoF的时候表示已经读取到了文件的末尾，跳出循环
			if err == io.EOF {
				fmt.Println("文件已经读取到了最后了")
				break
			}
			//如果不是那个特殊的错误则表示是在读取文件当中发生了错误，结束函数
			fmt.Println("文件读写异常err", err)
			return
		} else {
			fmt.Print(str)
		}
	}
	fmt.Println("文件读取完成")
}

//一次性文件读取
func main21() {
	bytes, err := ioutil.ReadFile("C:/Users/wuming/Desktop/读写文件测试.txt")
	if err == nil {
		fmt.Println(string(bytes))
	} else {
		fmt.Println("读取失败err,", err)
	}
}
//缓冲文件写出
func main22() {
	file, err := os.OpenFile("C:/Users/wuming/Desktop/读写文件测试.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	defer func() {
		file.Close()
		fmt.Println("文件已关闭")
	}()
	if err != nil {
		fmt.Println("文件打开失败err=", err)
		return
	}
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("\n我在里面添加了数据哈哈哈\n这是我添加数据的第二行")
	if err != nil {
		fmt.Println("写入缓冲区失败err,", err)
		return
	}
	_ = writer.Flush()
	fmt.Println("写入完成")
}
//一次性文件写出
func main23(){
	data := `
		静夜思
        --李白
	  床前明月光，
	  疑似地下霜。
	  举头望明月，
	  低头思故乡。
`
	fmt.Println(data)
	err := ioutil.WriteFile("C:/Users/wuming/Desktop/读写文件测试1.txt", []byte(data), 0666)
	if err != nil {
		fmt.Println("写入错误，err=",err)
	}else {
		fmt.Println("写出成功")
	}
}
func main24 (){
	fileInfo, err := os.Stat("C:/Users/wuming/Desktop/读写文件测试1.txt")
	if err != nil {
		if os.IsNotExist(err){
			fmt.Println("文件不存在的错误err=",err)
		}
		fmt.Println("err =",err)
	}else {
		fmt.Println(fileInfo)
	}
}
