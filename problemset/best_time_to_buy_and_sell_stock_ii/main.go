package main

import "fmt"

// Input: [7,1,5,3,6,4]
// Output: 7
// Explanation:
// Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
// Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
func maxProfit(prices []int) int {
	p := 0
	for i := 0; i < len(prices)-1; i++ {
		fmt.Println("target:", prices[i], prices[i+1])
		if prices[i+1] > prices[i] {
			fmt.Println("hit:", prices[i], prices[i+1])
			p += prices[i+1] - prices[i]
		}
	}
	return p
}
