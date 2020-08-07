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

// less space complexity
func minFallingPathSum(A [][]int) int {
	n, m := len(A), len(A[0])

	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			aij := A[i][j]
			A[i][j] = math.MaxInt32
			for jj := max(0, j-1); jj <= min(m-1, j+1); jj++ {
				A[i][j] = min(A[i][j], A[i-1][jj]+aij)
			}
		}
		// fmt.Println(dp[i])
	}

	ans := math.MaxInt32
	for j := range A[0] {
		ans = min(ans, A[n-1][j])
	}
	return ans
}

func minFallingPathSum0(A [][]int) int {
	n, m := len(A), len(A[0])

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			if i == 0 {
				dp[i][j] = A[i][j]
				continue
			}
			dp[i][j] = math.MaxInt32
		}
	}

	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			for jj := max(0, j-1); jj <= min(m-1, j+1); jj++ {
				dp[i][j] = min(dp[i][j], dp[i-1][jj]+A[i][j])
			}
		}
		// fmt.Println(dp[i])
	}

	ans := math.MaxInt32
	for j := range dp[0] {
		ans = min(ans, dp[n-1][j])
	}
	return ans
}

func main() {
	fmt.Println(minFallingPathSum([][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}))

	fmt.Println(minFallingPathSum([][]int{
		[]int{-80, -13, 22},
		[]int{83, 94, -5},
		[]int{73, -48, 61},
	}))
}
