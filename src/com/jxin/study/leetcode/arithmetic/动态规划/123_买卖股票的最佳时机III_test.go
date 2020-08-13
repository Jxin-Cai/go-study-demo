package 动态规划

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {

	maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4})
}

func maxProfit(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}
	k := 2

	pd := initPd(prices)
	const min = ^int(^uint(0) >> 1)
	pd[0][2][0] = 0
	pd[0][1][1] = -prices[0]
	pd[0][1][0] = 0
	pd[0][0][1] = -prices[0]
	pd[0][0][0] = 0

	pd[1][2][0] = 0
	pd[1][1][1] = max(pd[0][1][1], -prices[1])
	pd[1][1][0] = prices[1] - prices[0]
	pd[1][0][1] = max(pd[0][1][1], -prices[1])
	pd[1][0][0] = 0
	for i := k; i < len(prices); i++ {
		pd[i][k][0] = 0
		for j := 0; j < k; j++ {
			pd[i][j][0] = max(pd[i-1][j][0], pd[i-1][j][1]+prices[i])
			pd[i][j][1] = max(pd[i-1][j][1], pd[i-1][j+1][0]-prices[i])
		}
	}
	ret := 0
	for i := 0; i <= k; i++ {
		ret = max(ret, pd[len(prices)-1][i][0])
	}
	return ret
}

func initPd(prices []int) [][][]int {
	pd := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		pd[i] = make([][]int, 3)
		for j := 0; j < 3; j++ {
			pd[i][j] = make([]int, 2)
		}
	}

	return pd
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
