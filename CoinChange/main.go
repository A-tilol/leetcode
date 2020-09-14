package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
		for j := range coins {
			if i-coins[j] < 0 || dp[i-coins[j]] == -1 {
				continue
			}
			if dp[i] == -1 {
				dp[i] = dp[i-coins[j]] + 1
				continue
			}
			dp[i] = min(dp[i], dp[i-coins[j]]+1)
		}
	}

	return dp[amount]
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1, 2, 5, 15}, 11))
	fmt.Println(coinChange([]int{2, 5, 10, 1}, 27))
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
}
