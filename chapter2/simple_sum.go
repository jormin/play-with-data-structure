package chapter2

// 简单加法
func simpleSum(max int) (sum int) {
	for i := 1; i <= max; i++ {
		sum += i
	}
	return sum
}
