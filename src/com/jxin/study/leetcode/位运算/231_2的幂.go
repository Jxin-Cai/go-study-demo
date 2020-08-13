package 位运算

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	n &= n - 1
	return n == 0
}
