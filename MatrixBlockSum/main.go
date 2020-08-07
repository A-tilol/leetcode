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

// less time complexity: 2 dimension cumulative sum (ruisekiwa)
func matrixBlockSum(mat [][]int, K int) [][]int {
	n, m := len(mat), len(mat[0])
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}

	// ruisekiwa
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			mat[i][j] += mat[i][j-1]
		}
		if i == 0 {
			continue
		}
		for j := 0; j < m; j++ {
			mat[i][j] += mat[i-1][j]
		}
	}
	// fmt.Println(mat)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans[i][j] += mat[min(n-1, i+K)][min(m-1, j+K)]
			if i-K > 0 {
				ans[i][j] -= mat[i-K-1][min(m-1, j+K)]
			}
			if j-K > 0 {
				ans[i][j] -= mat[min(n-1, i+K)][j-K-1]
			}
			if i-K > 0 && j-K > 0 {
				ans[i][j] += mat[i-K-1][j-K-1]
			}
		}
	}

	return ans
}

// less time complexity: 1 dimension cumulative sum (ruisekiwa)
func matrixBlockSum1(mat [][]int, K int) [][]int {
	n, m := len(mat), len(mat[0])
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}

	// ruisekiwa of row
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			mat[i][j] += mat[i][j-1]
		}
		// fmt.Println(mat[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for r := max(0, i-K); r <= min(n-1, i+K); r++ {
				ans[i][j] += mat[r][min(m-1, j+K)]
				if j-K > 0 {
					ans[i][j] -= mat[r][j-K-1]
				}
			}
		}
	}

	return ans
}

func matrixBlockSum0(mat [][]int, K int) [][]int {
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
