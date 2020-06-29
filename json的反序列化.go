package main

import (
	"encoding/json"
	"fmt"
)


var jsonString =  `{"Name":"于谦","Age":50,"Rmb":12345.36,"Sex":true,"Hobby":["抽烟","喝酒","烫头"]}`
//json字符串反序列化成map
func main031()  {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &jsonMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonMap)
}
// json字符串反序列化成结构体
func main032(){
	type Person struct {
		Name  string
		Age   int
		Rmb   float64
		Sex   bool
		Hobby []string
	}
	 p := new(Person)
	err := json.Unmarshal([]byte(jsonString),  p)
	if err!= nil {
		panic(err)
	}
	fmt.Println(p)
}

// json字符串反序列化成切片
func main033(){

}