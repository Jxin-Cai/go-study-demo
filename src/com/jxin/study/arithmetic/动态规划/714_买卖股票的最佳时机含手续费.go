package 动态规划

func maxProfit(prices []int, fee int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}

	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0] - fee

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
	}
	l := len(prices)
	return dp[l-1][0]
}

func max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}
