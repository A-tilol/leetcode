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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func matrixBlockSum(mat [][]int, K int) [][]int {
	n, m := len(mat), len(mat[0])
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for r := max(0, i-K); r <= min(n-1, i+K); r++ {
				for c := max(0, j-K); c <= min(m-1, j+K); c++ {
					ans[i][j] += mat[r][c]
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(matrixBlockSum(
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		}, 1))

	fmt.Println(matrixBlockSum(
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		}, 2))
}
