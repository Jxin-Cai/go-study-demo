package arithmetic

import "testing"

func TestArr(t *testing.T) {

	longestValidParentheses("((()))())")
}
func longestValidParentheses(s string) int {
	var i []int
	var retArr []int

	left := 0
	temp := 0
	for _, v := range s {
		if string(v) == "(" {
			left++
			if left > 1 {
				i = append(i, temp)
				temp = 0
				left--
			}
			continue
		}
		left--
		if left < 0 {
			l := len(i)
			if l == 0 {
				retArr = append(retArr, temp)
				temp = 0
			} else {
				temp += i[l-1]
				i = i[:l-1]
				temp += 2
			}
			left++
			continue
		}
		temp += 2
	}
	retArr = append(retArr, temp)
	retArr = append(retArr, i...)
	ret := 0
	for _, v := range retArr {
		ret = max(ret, v)
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
