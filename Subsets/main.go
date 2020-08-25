package main

import "fmt"

var ns []int
var ans [][]int

func dfs(i, n int, a []int) {
	if i == n {
		ans = append(ans, a)
		return
	}
	acopy := make([]int, len(a))
	copy(acopy, a)
	dfs(i+1, n, a)
	dfs(i+1, n, append(acopy, ns[i]))
}

func subsets(nums []int) [][]int {
	ns = nums
	ans = make([][]int, 0)
	dfs(0, len(nums), []int{})
	return ans
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
