package 字符串

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	l := len(s)
	arr1, arr2 := []byte(s), []byte(t)
	cache1, cache2 := make(map[byte]byte), make(map[byte]byte)
	for i := 0; i < l; i++ {
		_, ok1 := cache1[arr1[i]]
		_, ok2 := cache2[arr2[i]]
		if ok1 != ok2 {
			return false
		}
		if cache1[arr1[i]] == arr2[i] {
			continue
		}
		cache1[arr1[i]] = arr2[i]
		cache2[arr2[i]] = arr1[i]
		return false
	}
	return true
}
