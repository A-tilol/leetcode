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

// less space complexity
func countSquares(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				continue
			}
			if i == 0 || j == 0 {
				ans++
				continue
			}
			matrix[i][j] = min(matrix[i-1][j-1], min(matrix[i-1][j], matrix[i][j-1])) + 1
			ans += matrix[i][j]
		}
		// fmt.Println(matrix[i])
	}
	return ans
}

// less time complexity
func countSquares1(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				continue
			}
			if i == 0 || j == 0 {
				dp[i][j] = 1
				ans++
				continue
			}
			dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			ans += dp[i][j]
		}
		// fmt.Println(dp[i])
	}
	return ans
}

// ruisekiwa
func countSquares0(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])

	// cumulative sum
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			matrix[i][j] += matrix[i][j-1]
		}
		if i == 0 {
			continue
		}
		for j := 0; j < m; j++ {
			matrix[i][j] += matrix[i-1][j]
		}
	}
	// fmt.Println(matrix)

	ans := 0
	for side := 1; side <= min(n, m); side++ {
		for i := side - 1; i < n; i++ {
			for j := side - 1; j < m; j++ {
				oneCountOfSquare := matrix[i][j]
				if i-side >= 0 {
					oneCountOfSquare -= matrix[i-side][j]
				}
				if j-side >= 0 {
					oneCountOfSquare -= matrix[i][j-side]
				}
				if i-side >= 0 && j-side >= 0 {
					oneCountOfSquare += matrix[i-side][j-side]
				}
				// fmt.Println(oneCountOfSquare, side*side)
				if oneCountOfSquare == side*side {
					ans++
				}
			}
		}
		// fmt.Println("ans:", ans)
	}
	return ans
}

func main() {
	fmt.Println(countSquares([][]int{
		[]int{0, 1, 1, 1},
		[]int{1, 1, 1, 1},
		[]int{0, 1, 1, 1},
	}) == 15)

	fmt.Println(countSquares([][]int{
		[]int{1, 0, 1},
		[]int{1, 1, 0},
		[]int{1, 1, 0},
	}) == 7)
}
