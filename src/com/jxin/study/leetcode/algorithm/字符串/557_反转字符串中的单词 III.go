package 字符串

import "strings"

func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}
	arr := strings.Fields(s)

	for i := 0; i < len(arr); i++ {
		arr[i] = reverseStr(arr[i])
	}
	return strings.Join(arr, " ")
}

func reverseStr(s string) string {
	if len(s) < 0 {
		return ""
	}
	arr := []byte(s)
	i := 0
	l := len(s) - 1
	for i < l {
		arr[i], arr[l] = arr[l], arr[i]
		i++
		l--
	}
	return string(arr)
}
