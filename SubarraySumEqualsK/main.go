package main

import (
	"fmt"
)

// O(n)
func subarraySum(nums []int, k int) int {
	m := map[int]int{0: 1}
	ans, sum := 0, 0
	for i := range nums {
		sum += nums[i]
		if v, ok := m[sum-k]; ok {
			ans += v
		}
		m[sum]++
	}
	return ans
}

// O(n^2)
func subarraySum1(nums []int, k int) int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	nums = append([]int{0}, nums...)

	ans := 0
	for i := range nums {
		for j := i + 1; j < n+1; j++ {
			if nums[j]-nums[i] == k {
				ans++
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 1, 1, 2, 1, 2}, 3))
}
