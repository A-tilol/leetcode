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

// less space complexity
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	pre0, pre1 := 0, nums[0]
	for i := 1; i < n; i++ {
		pre0, pre1 = max(pre0, pre1), pre0+nums[i]
	}

	return max(pre0, pre1)
}

func rob0(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0] = []int{0, nums[0]}
	// fmt.Println(dp[0])
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
		// fmt.Println(dp[i])
	}

	return max(dp[n-1][0], dp[n-1][1])
}

func main() {
	fmt.Println(rob([]int{}) == 0)
	fmt.Println(rob([]int{1}) == 1)
	fmt.Println(rob([]int{1, 2, 3, 1}) == 4)
	fmt.Println(rob([]int{2, 7, 9, 3, 1}) == 12)
}
