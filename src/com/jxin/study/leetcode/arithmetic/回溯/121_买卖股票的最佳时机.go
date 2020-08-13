package 回溯

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	ret := 0
	pre := prices[0]
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] > pre {
			continue
		}
		max := max(prices[i+1:])
		count := max - prices[i]
		if count > ret {
			ret = count
		}
		pre = prices[i]
	}

	return ret

}

func max(prices []int) int {
	ret := 0
	for _, v := range prices {
		if v <= ret {
			continue
		}
		ret = v
	}
	return ret
}
