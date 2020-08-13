package 动态规划

import "testing"

func TestMaxProfit(t *testing.T) {

	maxProfit([]int{1, 2, 3, 0, 2})
}

func maxProfit(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}

	pd := initPd(prices)
	pd[0][0] = 0
	pd[0][1] = -prices[0]

	pd[1][0] = max(0, prices[1]+pd[0][1])
	pd[1][1] = max(-prices[0], -prices[1])
	for i := 2; i < len(prices); i++ {
		pd[i][0] = max(pd[i-1][0], pd[i-1][1]+prices[i])
		pd[i][1] = max(pd[i-1][1], pd[i-2][0]-prices[i])

	}
	return pd[len(prices)-1][0]
}
func initPd(prices []int) [][]int {
	pd := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		pd[i] = make([]int, 2)
	}

	return pd
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
