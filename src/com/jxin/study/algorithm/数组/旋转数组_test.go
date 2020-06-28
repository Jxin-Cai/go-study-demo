package 数组

import "testing"

func TestService(t *testing.T) {
	rotate2([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
func rotate1(nums []int, k int) {
	l := len(nums)
	k %= l
	copy(nums, append(nums[l-k:], nums[0:l-k]...))
}
func rotate2(nums []int, k int) {
	k %= len(nums)
	reversal(nums)
	reversal(nums[0:k])
	reversal(nums[k:])
}
func reversal(arr []int) {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
	}
}
