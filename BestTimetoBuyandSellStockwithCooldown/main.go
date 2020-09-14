package main

import (
	"fmt"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 3)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
	}
	dp[0][0], dp[0][1], dp[0][2] = -prices[0], 0, 0

	for i := 1; i < n; i++ {
		// j==0:buy, j==1:sell, j==2:cooldown
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = max(dp[i-1][0]+prices[i], dp[i-1][1])
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}

	// fmt.Println(dp)

	return max(dp[n-1][0], max(dp[n-1][1], dp[n-1][2]))
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit([]int{1, 5, 2, 3, 0, 2}))
}
