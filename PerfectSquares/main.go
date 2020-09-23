package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var dp = make([]int, int(1e6))
var maxn = 1

func numSquares(n int) int {
	if n > len(dp) {
		dp = append(dp, make([]int, n-len(dp))...)
	}
	for i := maxn; i < n+1; i++ {
		dp[i] = math.MaxInt32
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	maxn = max(maxn, n+1)
	return dp[n]
}

func main() {
	fmt.Println(numSquares(1))
	fmt.Println(numSquares(2))
	fmt.Println(numSquares(4))
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
	fmt.Println(numSquares(6))
}
