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

var memo = make([][]int, 201)

func comb(n, i int) int {
	if n == 1 || i == 0 || i == n {
		return 1
	}
	if memo[n][i] != 0 {
		return memo[n][i]
	}
	memo[n][i] = comb(n-1, i-1) + comb(n-1, i)
	return memo[n][i]
}

func uniquePaths(m int, n int) int {
	for i := range memo {
		memo[i] = make([]int, 101)
	}
	return comb(m+n-2, min(m, n)-1)
}

func main() {
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(7, 3))
	fmt.Println(uniquePaths(51, 9))
}
