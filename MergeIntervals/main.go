package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0, len(intervals))
	if len(intervals) == 0 {
		return ans
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	t := intervals[0]
	for _, interval := range intervals[1:] {
		if t[1] >= interval[0] {
			if t[1] <= interval[1] {
				t[1] = interval[1]
			}
		} else {
			ans = append(ans, t)
			t = interval
		}
	}
	ans = append(ans, t)
	return ans
}

func main() {
	fmt.Println(merge([][]int{}))
	fmt.Println(merge([][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}))
	fmt.Println(merge([][]int{
		[]int{1, 4},
		[]int{4, 5},
	}))
	fmt.Println(merge([][]int{
		[]int{1, 4},
		[]int{0, 4},
	}))
	fmt.Println(merge([][]int{
		[]int{1, 4},
		[]int{2, 3},
	}))
}
