package main

func main() {
	//	myRand := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	//	answer := myRand.Intn(1000)
	//	int
	//loop:
	//	for {
	//		//接受用户输入
	//		fmt.Println("请输入您的猜想数字:")
	//		var guess string
	//		fmt.Scan(&guess)
	//		guessNum, _ := strconv.Atoi(guess)
	//		if guess == "off"{
	//			break
	//		}
	//		switch {
	//		case guessNum > answer:
	//			fmt.Println("猜大了")
	//		case guessNum < answer:
	//			fmt.Println("猜小了")
	//		default:
	//			fmt.Println("猜对了")
	//			break loop
	//		}
	//	}

	//n1 := 123
	//
	//fmt.Printf("n1的数据类型是%T", n1)
	//
	//fmt.Println()

	//fmt.Println(180/(math.Pi/math.Asin(0.5)))
	//for i:=0;i<100;i++{
	//	fmt.Println(i,"\t",fib(i))
	//}
	//dir,err :=os.Getwd()
	//fmt.Println(dir)
	//fmt.Println(err)

	//获得指定的环境变量
	//goroot := os.Getenv("GOROOT")
	//gopath := os.Getenv("GOPATH")
	//fmt.Println(gopath)
	//fmt.Println(goroot)

	//获取所有环境变量
	//envs := os.Environ()
	//for _,env := range envs{
	//	fmt.Println(env)
	//}

	//获取主机名
	//host,err := os.Hostname()
	//if err ==nil{
	//	fmt.Println(host)
	//}

	//判断系统分隔符
	//fmt.Println("/是系统分隔符吗",os.IsPathSeparator('/'))
	//fmt.Println("\\是系统分隔符吗",os.IsPathSeparator('\\'))
	//fmt.Println("$是系统分隔符吗",os.IsPathSeparator('$'))


}
func fib(n int) int{
	if n==0 || n==1{
		return 1
	}
	return fib(n-1)+fib(n-2)
}
