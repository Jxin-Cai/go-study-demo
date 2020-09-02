package 动态规划

func numDecodings(s string) int {
	if len(s) < 1 {
		return 0
	}
	dp := make([]int, len(s))
	if s[0] == '0' {
		return 0
	}
	dp[0] = 1
	if len(s) < 2 {
		return dp[len(s)-1]
	}

	if s[1] == '0' {
		if s[0] != '1' && s[0] != '2' {
			return 0
		}
		dp[1] = dp[0]
	} else {
		if s[0] == '1' {
			dp[1] = 2
		} else if s[0] == '2' && s[1] <= '6' {
			dp[1] = 2
		} else {
			dp[1] = dp[0]
		}
	}

	for i := 2; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] != '1' && s[i-1] != '2' {
				return 0
			}
			dp[i] = dp[i-2]
			continue
		}

		if s[i-1] == '1' {
			dp[i] = dp[i-1] + dp[i-2]
			continue
		}
		if s[i-1] == '2' && s[i] <= '6' {
			dp[i] = dp[i-1] + dp[i-2]
			continue
		}
		dp[i] = dp[i-1]
	}
	return dp[len(s)-1]
}
