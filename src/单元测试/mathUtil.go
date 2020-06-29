package mymath

// 用循环的方式求自然数的和
func GetSum(n int) (sum int) {
	for i := 0; i <= n; i++ {
		sum += i
	}
	return
}

//用递归的方法求自然数的和
func GetSumRecursively(n int) (sum int) {
	if n == 1 {
		return 1
	}
	return n + GetSumRecursively(n-1)
}
