package 字符串

import (
	"strings"
)

func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}
	arr := strings.Fields(s)
	i := 0
	l := len(arr) - 1
	for i < l {
		arr[i], arr[l] = arr[l], arr[i]
		i++
		l--
	}
	return strings.Join(arr, " ")
}
