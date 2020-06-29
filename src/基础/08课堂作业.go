package main

import (
	"fmt"
	"mylib"
)

func main() {
	//i := []int{1, 2, 8, 6, 3, 4, 5, 0, 9}
	//fmt.Println("排序前", i)
	//sortSlice(i)
	//fmt.Println("排序后", i)
	// 班内的学生切片
	student := []string{"小铭", "小花", "小华", "小朱", "小陈", "小美", "小胖", "大雄", "静香", "胖虎", "小夫"}
	studentScoreMap := mylib.StudentScore(student)
	sortKey := mylib.SoreStudentScore(studentScoreMap)
	fmt.Println("成绩单")
	for _, v := range sortKey {
		fmt.Printf("姓名：%s,成绩：%d\n", v, studentScoreMap[v])
	}
}

//排序函数
func sortSlice(slice []int) {
	var temp int
	lens := len(slice)
	for i := 0; i < lens-1; i++ {
		for j := 0; j < lens-1-i; j++ {
			if slice[j] < slice[j+1] {
				temp = slice[j+1]
				slice[j+1] = slice[j]
				slice[j] = temp
			}
		}
	}
}
