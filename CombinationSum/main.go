package main

import (
	"fmt"
	"sort"
)

var ans [][]int
var numbers []int
var tar int

func dfs(i int, comb []int, sum int) {
	if sum == tar {
		ans = append(ans, comb)
		return
	}
	if i == len(numbers) {
		return
	}
	if sum+numbers[i] > tar {
		return
	}

	combCopy := make([]int, len(comb))
	copy(combCopy, comb)

	comb = append(comb, numbers[i])
	dfs(i, comb, sum+numbers[i])

	dfs(i+1, combCopy, sum)
}

func combinationSum(candidates []int, target int) [][]int {
	ans = make([][]int, 0)
	numbers = candidates
	sort.Ints(numbers)
	tar = target
	dfs(0, []int{}, 0)
	return ans
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
