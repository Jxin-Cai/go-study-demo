package 字符串

func firstUniqChar(s string) int {
	if s == "" {
		return -1
	}
	cache := make(map[int32]int)
	for _, v := range s {
		if _, ok := cache[v]; ok {
			cache[v] = -1
			continue
		}
		cache[v] = 1
	}

	for i, v := range s {
		if cache[v] == 1 {
			return i
		}
	}
	return -1
}
