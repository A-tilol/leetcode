package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	start := l

	if nums[start] != target {
		return []int{-1, -1}
	}

	l, r = start, len(nums)-1
	for l < r {
		m := (l+r)/2 + 1
		if nums[m] <= target {
			l = m
		} else {
			r = m - 1
		}
	}

	return []int{start, l}
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 0))
	fmt.Println(searchRange([]int{1}, 1))
}
