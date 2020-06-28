package 数组

func moveZeroes1(nums []int) {
	if nums == nil {
		return
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}

		if i == j {
			j++
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		j++
	}
}
func moveZeroes2(nums []int) {
	if nums == nil {
		return
	}
	i := 0
	for i < len(nums) && nums[i] != 0 {
		i++
	}
	j := i
	for ; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		j++
	}
}
func moveZeroes3(nums []int) {
	// 递减非零计数器
	l := len(nums)
	// 递增非零计数器
	i := 0
	for i < l {
		if nums[i] != 0 {
			i++
			continue
		}
		l--
		// 把0前后两部分合并
		nums = append(nums[0:i], nums[i+1:]...)
		// 在末尾补回0
		nums = append(nums, 0)
	}
}
