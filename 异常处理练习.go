package main

import (
	"fmt"
)

/**
√ 封装人，属性包括姓名、年龄、身高、体重、颜值、资产、性别、性取向
√ 给人封装结婚方法，参数是潜在的结婚对象：
①如果对方的性取向有问题，panic
②如果对方的颜值过低，返回错误
③否则返回满意程度
 */

type Gender int

func String(g Gender) string {
	return []string{"Male", "Female", "Bisexual"}[g]
}

const (
	Male     = iota // 男的
	Female          // 女的
	Bisexual        // 二椅子
)

type BadSpouseError struct {
	why string
}

func (bse *BadSpouseError) Error() string {
	return bse.why
}

func (bse *BadSpouseError) String() string {
	return bse.why
}

func CreateBadSpouseError(o *Human) *BadSpouseError {
	if o.Rmb < 1e6 {
		return &BadSpouseError{"太穷了，看不上"}
	}

	if o.Weight > 200 {
		return &BadSpouseError{"太胖了，驼不住"}
	}

	if o.Age > 50 {
		return &BadSpouseError{"太大了，不合适"}
	}

	return nil
}

type Human struct {
	Name          string // 性别
	Age           int    // 年龄
	Height        int    // 身高
	Weight        int    //体重
	Looking       int    // 颜值
	TargetLooking int    // 对象的颜值
	Rmb           int    //人民币
	Sex           Gender //自己的性别
	TargetSex     Gender //对象的性别
}

func (h *Human) Marry(o *Human) (happiness int, err error) {

	//如果对方性别和自己的目标性别不符合恐慌
	if o.Sex != h.TargetSex {
		panic(&BadSpouseError{"性别不符合要求"})
	}
	if err = CreateBadSpouseError(o); err != nil {
		return
	}
	//计算幸福程度
	happiness = o.Height * o.Rmb * o.Looking / o.Age * o.Weight
	return

}

func NewHuman(name string, age, height, weight, looking, targetLooking, rmb int, sex, targetSex Gender) *Human {
	human := Human{}
	human.Name = name
	human.Age, human.Height, human.Weight, human.Looking, human.TargetLooking, human.Rmb = age, height, weight, looking, targetLooking, rmb
	human.Sex, human.TargetSex = sex, targetSex
	return &human
}
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	cook := NewHuman("库克", 60, 180, 150, 60, 90, 1.2e10, Male, Male)
	ySister := NewHuman("你妹", 20, 180, 150, 99, 90, 1.2e10, Female, Male)
	//happiness, err := cook.Marry(ySister)
	happiness, err := ySister.Marry(cook)
	if err != nil {
		fmt.Println("牵手失败,err=", err)
	} else {
		fmt.Println("牵手成功,幸福指数=", happiness)
	}
}
