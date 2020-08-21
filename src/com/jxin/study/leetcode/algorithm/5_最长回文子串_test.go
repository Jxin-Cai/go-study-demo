package algorithm

import "testing"

func TestLong(t *testing.T) {
	longestPalindrome("bb")
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	l := len(s)
	dp := make([][]bool, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
	}
	ret := ""
	for distance := 0; distance < l; distance++ {
		i := 0
		j := i + distance
		for j < l {
			if distance <= 2 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
			if dp[i][j] && j-i+1 > len(ret) {
				ret = s[i : j+1]
			}
			i++
			j++
		}
	}
	return ret
}
