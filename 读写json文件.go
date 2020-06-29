package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main041(){
	file, err := os.OpenFile("F:/golandWoke/file/于谦.json", os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer func (){
		file.Close()
		fmt.Println("文件已关闭")
	}()
	dataMap := make(map[string]interface{})
	dataMap["Name"] = "于谦"
	dataMap["Age"] = 50
	dataMap["Rmb"] = 12345.36
	dataMap["Sex"] = true
	dataMap["Hobby"] = []string{"抽烟", "喝酒", "烫头"}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(dataMap)
	if err != nil {
		panic(err)
	}
	fmt.Println("文件写入成功")
}

func main (){
	jsonMap := make(map[string] interface{})
	file, err := os.OpenFile("F:/golandWoke/file/于谦.json", os.O_RDONLY, 0666)
	if err!= nil {
		panic(err)
	}
	defer func (){
		file.Close()
		fmt.Println("关闭成功")
	}()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&jsonMap)
	if err != nil {
		panic(err)
	}
	fmt.Println("解码成功:",jsonMap)
}
