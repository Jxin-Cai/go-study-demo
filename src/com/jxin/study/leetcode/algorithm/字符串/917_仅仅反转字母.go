package 字符串

func reverseOnlyLetters(S string) string {
	if len(S) == 0 {
		return S
	}
	arr := []byte(S)

	i := 0
	l := len(arr) - 1

	for i < l {
		if !isLetter(arr[i]) {
			i++
			continue
		}
		if !isLetter(arr[l]) {
			l--
			continue
		}
		arr[i], arr[l] = arr[l], arr[i]
		i++
		l--
	}
	return string(arr)
}

func isLetter(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}
