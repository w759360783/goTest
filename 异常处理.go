package main

import (
	"errors"
	"fmt"
	"math"
)

func main022() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	ret, err := Area(-10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("面积为：",ret)
}

func Area(r float64) (float64, error) {
	switch {
	case r < 0:
		panic("半径不能为负数")
	case r < 5 || r > 50:
		return 0.0, errors.New("不合规范的参数")
	default:
		return r * r * math.Pi, nil
	}
}
