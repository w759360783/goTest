package main

import "fmt"

//最简单的获得命令行参数

func main09() {
	//for i, v := range os.Args {
	//	fmt.Println(i, v)
	//}
}

func main10() {
	////-name 吴铭 -age 23 -sex 男 -isSmoke false
	//
	//namePtr := flag.String("name", "无名氏", "姓名")
	//agePtr := flag.Int("age", -1, "年龄")
	//sexPtr := flag.String("sex", "未填写", "性别")
	//isSmoke := flag.Bool("isSmoke", true, "是否抽烟")
	//flag.Parse()
	//fmt.Println(*namePtr, *agePtr, *sexPtr, *isSmoke)

}

type option interface {
	a()
}

type st struct {
	name string
}

func (s *st) a() {
	fmt.Println("a")
}
func (s *ss) a() {
	fmt.Println("a")
}

type ss struct {
	name string
}

func main() {
	st := st{"a"}
	o := option(&st)
	fmt.Println(o)
	fmt.Println(&st)
}
