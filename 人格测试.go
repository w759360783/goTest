package main

import (
	"flag"
	"fmt"
	"unicode/utf8"
)

/*
人格测试的需求:
从命令行输入如下信息:
renge.exe -name xxx -xingzuo xxx -age 40 -rmb 1234 -isfool true
Keys:
√人格与资产成正比，与年龄成反比
√姓名如果如果在四个字以上分数*=1.1
√觉得自己是个蠢猪的，人格成绩*=1.10
√年纪一把鸟钱没有、还觉得自己很聪明的，人格成绩*=0.8
√星座：天蝎、摩羯、水瓶 *=1.05，处女*=0.9
√人格系数：以40岁/40万=1
*/
func main() {
	score := GetRenge(GetCmdArgs())
	fmt.Printf("您的评分是%.2f\n",score)
	switch  {
	case score > 90.0:
		fmt.Println("膜拜大神,666")
	case score >60.0:
		fmt.Println("很普通")
	default:
		fmt.Println("垃圾")
	}

}

//评分计算系统
func GetRenge(name, xingzuo string, age, rmb float64, isfool bool) (score float64) {
	//进入评分系统
	//假设平均人格为60
	score = 60.0

	//如果姓名在四个字以上
	if utf8.RuneCountInString(name) >= 4 {
		score *= 1.1
	}
	//星座加分项（星座：天蝎、摩羯、水瓶 *=1.05，处女*=0.95）
	switch xingzuo {
	case "天蝎", "摩羯", "水瓶":
		score *= 1.1
	case "处女":
		score *= 0.9
	default:
	}
	//觉得自己是个蠢猪的加分
	if isfool {
		score *= 1.1
	}
	//年纪一把鸟钱没有、还觉得自己很聪明的，人格成绩*=0.8
	if age > 30 && rmb < 400000 && !isfool {
		score *= 0.8
	}
	//人格与资产成正比，与年龄成反比
	k := 1 / (400000.0 / 40.0)
	score *= k * rmb / age

	return
}

func GetCmdArgs() (name, xingzuo string, age, rmb float64, isfool bool) {
	namePrt := flag.String("name", "无名氏", "测试者的姓名")
	xingzuoPrt := flag.String("xingzuo", "没有星座", "测试者的星座")
	agePrt := flag.Float64("age", 0, "测试者的星座")
	rmbPrt := flag.Float64("rmb", 0, "测试者的星座")
	isfoolPrt := flag.Bool("isfool", true, "智商")
	flag.Parse()
	return *namePrt, *xingzuoPrt, *agePrt, *rmbPrt, *isfoolPrt
}
