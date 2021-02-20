package chapter2

import "testing"

// 最大数
var max = 100000000

// 期望结果
var wantSum = 5000000050000000

// 简单加法性能测试
func BenchmarkSimpleSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if gotSum := simpleSum(max); gotSum != wantSum {
			b.Errorf("gaussSum() = %v, want %v", gotSum, wantSum)
		}
	}
}

// 高斯算法性能测试
func BenchmarkGaussSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if gotSum := gaussSum(max); gotSum != wantSum {
			b.Errorf("gaussSum() = %v, want %v", gotSum, wantSum)
		}
	}
}
