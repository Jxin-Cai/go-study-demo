package 数组

import (
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {

	t.Log(threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}))
}
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	var ret [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i-1 >= 0 && nums[i] == nums[i-1] {
			continue
		}
		val := -nums[i]
		end := len(nums) - 1
		for j := i + 1; j < end; j++ {
			if j-2 >= i && nums[j] == nums[j-1] {
				continue
			}
			for nums[end]+nums[j] > val && end-1 > j {
				end--
			}

			if nums[end]+nums[j] == val {
				ret = append(ret, []int{nums[i], nums[j], nums[end]})
			}
		}
	}
	return ret
}
