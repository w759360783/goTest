package main

import "fmt"

func main() {

	const pi = 3.14
	var r = 9.0

	area := pi * r * r

	fmt.Println("圆的面积是：", area)

	r = 2.1
	area = pi * r * r
	fmt.Println("圆的面积是：", area)
}
