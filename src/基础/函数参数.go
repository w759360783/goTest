package main

import "fmt"

func main() {
	//sayHello()
	//sayHello1("小猪")
	//sayHello2("小猪", 10)
	//sayHello3("小猪","small_pig")
	sayHello4("小铭","小猪","小红","小华")
}

func sayHello() {
	fmt.Println("今天早上起床来")
	fmt.Println("你好呀(*^▽^*)")
}

func sayHello1(name string) {
	fmt.Println("今天早上起床来")
	fmt.Printf("%s你好呀", name)
}

func sayHello2(name string, n int) {
	fmt.Println("今天早上起床来")
	for i := 0; i < n; i++ {
		fmt.Printf("%s你好呀\n", name)
	}
}

func sayHello3(name, name1 string) {
	fmt.Println("今天早上起床来")
	fmt.Printf("%s你好呀\n", name)
	fmt.Printf("%s你好呀\n", name1)
}

func sayHello4(names ...string) {
	fmt.Println("今天早上起床来")
	for i, name := range names {
		fmt.Printf("No.%d,%s你好呀\n", i, name)
	}
}
