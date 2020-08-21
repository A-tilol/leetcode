package main

import (
	"fmt"
)

var ans [][]int

func dfs(nums []int, arr []int) {
	if len(nums) == 0 {
		ans = append(ans, arr)
		return
	}

	for i := range nums {
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)

		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)

		dfs(append(numsCopy[:i], numsCopy[i+1:]...), append(arr, nums[i]))
	}
}

func permute(nums []int) [][]int {
	ans = make([][]int, 0)
	dfs(nums, []int{})
	return ans
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
