package 数组

import (
	"testing"
)

func Test(t *testing.T) {
	removeDuplicates1([]int{1, 1, 2})
}

func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	v := nums[0]
	for idx, val := range nums {
		if val == v {
			continue
		}
		v = val
		j++
		nums[j] = nums[idx]
	}
	return j + 1
}
func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}

		nums[j] = nums[i]
		j++
	}
	return j
}
