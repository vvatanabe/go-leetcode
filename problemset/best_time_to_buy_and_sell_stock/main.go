package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var max int
	for i := 0; i < len(prices); i++ {
		for j := i; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if max < profit {
				max = profit
			}
		}
	}
	return max
}
