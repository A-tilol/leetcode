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

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	ans := math.MinInt32
	for i := range dp {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			}
			dp[i] = max(dp[i], dp[j]+1)
		}
		ans = max(ans, dp[i])
		// fmt.Println(dp)
	}
	return ans
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 4}))
	fmt.Println(lengthOfLIS([]int{5, 6, 7, 8, 4, 3, 2, 1}))
}
