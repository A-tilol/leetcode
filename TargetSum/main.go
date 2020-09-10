package main

import (
	"fmt"
	"sort"
)

var ns []int
var n, s int
var memo []map[int]int

func dfs(i, sum int) int {
	if i == n {
		if sum == s {
			return 1
		}
		return 0
	}
	if _, ok := memo[i][sum]; ok {
		return memo[i][sum]
	}

	pcnt := dfs(i+1, sum+ns[i])
	mcnt := dfs(i+1, sum-ns[i])

	memo[i][sum] = pcnt + mcnt
	return memo[i][sum]
}

func findTargetSumWays(nums []int, S int) int {
	sort.Ints(nums)
	ns, n, s = nums, len(nums), S

	memo = make([]map[int]int, n)
	for i := range memo {
		memo[i] = make(map[int]int)
	}

	return dfs(0, 0)
}

func main() {
	fmt.Println(findTargetSumWays(
		[]int{1, 1, 1, 1, 1}, 3,
	))
}
