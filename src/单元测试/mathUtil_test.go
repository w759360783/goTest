package mymath

import "testing"

func TestGetSum(t *testing.T) {
	sum := GetSum(10)

	if sum != 55 {
		t.Errorf("测试失败，期望结果50输出结果%d", sum)
		t.FailNow()
	}
	t.Log("测试成功，期望结果50输出结果", sum)
}

func TestGetSumRecursively(t *testing.T) {
	sum := GetSumRecursively(10)
	if sum != 55{
		t.Fatalf("测试失败，期望结果50，输出结果%d",sum)
	}
	t.Logf("测试成功，期望结果50，输出结果%d",sum)
}