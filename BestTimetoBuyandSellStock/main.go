package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProfit0(prices []int) int {
	ans := 0
	if len(prices) == 0 {
		return ans
	}

	hold := prices[0]
	for i := 1; i < len(prices); i++ {
		ans = max(ans, prices[i]-hold)
		hold = min(hold, prices[i])
	}
	return ans
}

func maxProfit(prices []int) int {
	ans := 0
	if len(prices) == 0 {
		return ans
	}

	hold := prices[0]
	for i := 1; i < len(prices); i++ {
		if ans < prices[i]-hold {
			ans = prices[i] - hold
		}
		if hold > prices[i] {
			hold = prices[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
