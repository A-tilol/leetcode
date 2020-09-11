package main

import (
	"fmt"
	"sort"
)

type A struct {
	num int
	n   int
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for i := range nums {
		m[nums[i]]++
	}

	as := make([]A, 0, len(m))
	for k, v := range m {
		as = append(as, A{k, v})
	}

	sort.Slice(as, func(i, j int) bool { return as[i].n > as[j].n })

	ans := make([]int, k)
	for i := range ans {
		ans[i] = as[i].num
	}

	return ans
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
}
