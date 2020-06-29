package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("当前协程的可用CPU核心数为", runtime.NumCPU())
	gomaxprocs := runtime.GOMAXPROCS(2)
	fmt.Println("先前的配置", gomaxprocs)
}
