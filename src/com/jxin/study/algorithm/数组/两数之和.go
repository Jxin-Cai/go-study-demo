package 数组

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for i, v := range nums {
		if idx, exist := cache[target-v]; exist {
			return []int{idx, i}
		}
		cache[v] = i
	}
	return nil
}
