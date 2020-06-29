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

func EncodePerson2JsonFile(p *Person, dstFile string) bool{
	file, err := os.OpenFile(dstFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败")
		return false
	}
	decoder := json.NewEncoder(file)
	err = decoder.Encode(p)
	if err != nil {
		fmt.Println("编码失败",err)
		return false
	} else {
		fmt.Println("编码成功")
		return true
	}
}

func DecodeJsonFile2Person(filePath string) (*Person, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDONLY|os.O_APPEND, 0666)
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
