package mymath

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string
	Age   int
	Rmb   int
	Sex   bool
	Hobby []string
}

func EncodePerson2JsonFile(p *Person, dstFile string) {
	file, err := os.OpenFile(dstFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败")
	}
	decoder := json.NewEncoder(file)
	err = decoder.Encode(p)
	if err != nil {
		fmt.Println("解码失败")
	} else {
		//fmt.Println("解码成功")
	}
}

func DecodeJsonFile2Person(filePath string) (*Person, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	person := new(Person)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(person)
	if err != nil {
		return nil, err
	}
	return person, nil
}
