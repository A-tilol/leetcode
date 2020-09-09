package main

import (
	"fmt"
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	s, e := 0, 0
	mini, maxi := n-1, 0
	for i := 0; i < n; i++ {
		if nums[i] >= nums[maxi] {
			maxi = i
		}
		if nums[n-1-i] <= nums[mini] {
			mini = n - 1 - i
		}
		if i != maxi {
			e = i
		}
		if n-1-i != mini {
			s = n - 1 - i
		}
	}

	if s == e {
		return 0
	}
	return e - s + 1
}

type array struct {
	i int
	n int
}

func findUnsortedSubarray1(nums []int) int {
	a := make([]array, len(nums))
	for i := range nums {
		a[i] = array{i, nums[i]}
	}
	sort.SliceStable(a, func(i, j int) bool { return a[i].n < a[j].n })

	s, e := -1, 0
	for i := range a {
		if a[i].i != i {
			if s == -1 {
				s = i
			}
			e = i
		}
	}

	if e == 0 {
		return 0
	}
	return e - s + 1
}

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println(findUnsortedSubarray([]int{1, 3, 2, 2, 2}))
	fmt.Println(findUnsortedSubarray([]int{2, 3, 3, 2, 4}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 4, 5, 3}))
}
