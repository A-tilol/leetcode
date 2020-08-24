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

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]

}

func minPathSum0(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	dp := make([][]int, m)
	done := make([][]bool, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
		done[i] = make([]bool, n)
	}
	dp[0][0] = grid[0][0]

	ps := make(map[int]map[int]struct{})
	ps[0] = map[int]struct{}{0: struct{}{}}

	for len(ps) > 0 {
		// fmt.Println(ps)

		x, y := m-1, n-1
		for px := range ps {
			for py := range ps[px] {
				if dp[px][py] < dp[x][y] {
					x, y = px, py
				}
			}
		}
		// fmt.Println("min", x, y)
		delete(ps[x],
			y)
		if len(ps[x]) == 0 {
			delete(ps, x)
		}
		done[x][y] = true

		if x == m-1 && y == n-1 {
			break
		}

		if x+1 < m && !done[x+1][y] {
			dp[x+1][y] = min(dp[x+1][y],
				dp[x][y]+grid[x+1][y])
			if _, ok := ps[x+1]; !ok {
				ps[x+1] = make(map[int]struct{})
			}
			ps[x+1][y] = struct{}{}
		}
		if y+1 < n && !done[x][y+1] {
			dp[x][y+1] = min(dp[x][y+1], dp[x][y]+grid[x][y+1])
			if _, ok := ps[x]; !ok {
				ps[x] = make(map[int]struct{})
			}
			ps[x][y+1] = struct{}{}
		}
	}

	// for i := range dp {
	// 	fmt.Println(dp[i])
	// }

	return dp[m-1][n-1]
}

func main() {
	fmt.Println(minPathSum([][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}))
	fmt.Println(minPathSum([][]int{
		[]int{1, 2, 5},
		[]int{3, 2, 1},
	}))
	fmt.Println(minPathSum([][]int{
		[]int{1, 4, 8, 6, 2, 2, 1, 7},
		[]int{4, 7, 3, 1, 4, 5, 5, 1},
		[]int{8, 8, 2, 1, 1, 8, 0, 1},
		[]int{8, 9, 2, 9, 8, 0, 8, 9},
		[]int{5, 7, 5, 7, 1, 8, 5, 5},
		[]int{7, 0, 9, 4, 5, 6, 5, 6},
		[]int{4, 9, 9, 7, 9, 1, 9, 0},
	}))
}
