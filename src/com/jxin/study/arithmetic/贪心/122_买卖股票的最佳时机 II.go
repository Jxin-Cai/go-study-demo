package 贪心

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	ret := 0
	pre := prices[0]
	for _, v := range prices {
		if v > pre {
			ret += v - pre
		}
		pre = v
	}
	return ret
}
