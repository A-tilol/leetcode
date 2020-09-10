package main

import "fmt"

func canPartition(nums []int) bool {
	n := len(nums)
	total := 0
	for i := range nums {
		total += nums[i]
	}
	if total%2 == 1 {
		return false
	}

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, total/2+1)
	}
	dp[0][0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j <= total/2; j++ {
			dp[i][j] = dp[i-1][j] || (j-nums[i-1] >= 0 && dp[i-1][j-nums[i-1]])
		}
		if dp[i][total/2] {
			return true
		}
	}
	return false
}

var n, total int
var ns []int
var memo []map[int]bool

func dfs(i, sum int) bool {
	if i == n {
		// fmt.Println(sum)
		return sum*2 == total
	}
	if _, ok := memo[i][sum]; ok {
		return false
	}

	if dfs(i+1, sum+ns[i]) {
		return true
	}
	if dfs(i+1, sum) {
		return true
	}

	memo[i][sum] = false
	return false
}

func canPartition1(nums []int) bool {
	n, ns = len(nums), nums
	total = 0
	for i := range nums {
		total += nums[i]
	}

	if total%2 == 1 {
		return false
	}

	memo = make([]map[int]bool, n)
	for i := range memo {
		memo[i] = make(map[int]bool)
	}

	return dfs(0, 0)
}

func main() {
	fmt.Println(canPartition(
		[]int{1, 5, 11, 5},
	))
	fmt.Println(canPartition(
		[]int{1, 2, 3, 5},
	))
}
