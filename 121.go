package main

func maxProfit(prices []int) int {
	ans := 0
	max := 0
	for i := len(prices) - 1; i >= 0; i-- {
		if max < prices[i] {
			max = prices[i]
		}
		if ans < max-prices[i] {
			ans = max - prices[i]
		}
	}
	return ans
}
