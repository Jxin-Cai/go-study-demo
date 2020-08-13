package 动态规划

import "testing"

func TestProfit(t *testing.T) {

	maxProfit(2, []int{2, 4, 1})
}
func maxProfit(k int, prices []int) int {
	if prices == nil || len(prices) < 2 || k == 0 {
		return 0
	}
	if k > len(prices)/2 {
		return maxProfit1(prices)
	}
	pd := initPd(prices, k)

	for i := 0; i < k; i++ {
		pd[0][i][0] = 0
		pd[0][i][1] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
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
func maxProfit1(prices []int) int {

	pd := initPd1(prices)
	pd[0][0] = 0
	pd[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		pd[i][0] = max(pd[i-1][0], pd[i-1][1]+prices[i])
		pd[i][1] = max(pd[i-1][1], pd[i-1][0]-prices[i])

	}
	return pd[len(prices)-1][0]
}
func initPd(prices []int, k int) [][][]int {
	pd := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		pd[i] = make([][]int, k+1)
		for j := 0; j < k+1; j++ {
			pd[i][j] = make([]int, 2)
		}
	}

	return pd
}
func initPd1(prices []int) [][]int {
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
