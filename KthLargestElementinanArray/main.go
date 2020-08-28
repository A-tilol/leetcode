package main

import (
	"fmt"
	"math"
	"sort"
)

func quickSelect(a []int, start, end, k int) int {
	if start == end {
		return a[start]
	}

	pivot := a[start]
	l, r := start, end
	for l < r {
		for l < r && a[r] >= pivot {
			r--
		}
		for l < r && a[l] < pivot {
			l++
		}
		a[l], a[r] = a[r], a[l]
	}

	if l >= len(a)-k {
		return quickSelect(a, start, l, k)
	}
	return quickSelect(a, l+1, end, k)
}

// algorithm: quick select
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k)
}

// O(nlogn)
func findKthLargest2(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// O(n*k)
func findKthLargest1(nums []int, k int) int {
	ans := make([]int, k)
	for j, n := range nums {
		if j < k {
			ans[j] = math.MinInt32
		}
		for i := 0; i < k; i++ {
			if n > ans[i] {
				ans[i], n = n, ans[i]
			}
		}
	}
	// fmt.Println(ans)
	return ans[k-1]
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	fmt.Println(findKthLargest([]int{5, 6, 5, 6}, 2))
}
