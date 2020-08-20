package 字符串

func validPalindrome(s string) bool {
	i, l := 0, len(s)-1
	for i < l {
		if s[i] != s[l] {
			return isPalindrome(i+1, l, s) || isPalindrome(i, l-1, s)
		}
		i++
		l--
	}
	return true
}

func isPalindrome(i int, l int, s string) bool {
	for i < l {
		if s[i] != s[l] {
			return false
		}
		i++
		l--
	}
	return true
}
