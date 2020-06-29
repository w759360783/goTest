package mymath

import "testing"

func BenchmarkGetSum(b *testing.B) {
	b.Log("BenchmarkGetSum测试开始")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetSum(10)
	}
}

func BenchmarkGetSumRecursively(b *testing.B) {
	b.Log("BenchmarkGetSumRecursively测试开始")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetSumRecursively(10)
	}
}

func BenchmarkGetSum2(b *testing.B) {
	b.Log("BenchmarkGetSum2测试开始")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetSum(100)
	}
}

func BenchmarkGetSumRecursively2(b *testing.B) {
	b.Log("BenchmarkGetSumRecursively2测试开始")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetSumRecursively(100)
	}
}
