package main

import "fmt"

func main() {
	//result := sum(15, 16)
	//fmt.Println("result=", result)

	he, cha, ji, shang := calc(15, 16)
	fmt.Println(he, cha, ji, shang)
}

func sum(num1, num2 int) (result int) {
	result = num1 + num2
	return
}

func calc(num1, num2 int) (he, cha, ji, shang int) {
	he = num1 + num2
	cha = num1 - num2
	ji = num1 * num2
	shang = num1 / num2
	return
}
