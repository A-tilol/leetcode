package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// less space complexity
func maxSubArray(nums []int) int {
	rui := 0
	ans, ruiMin := nums[0], 0
	for i := 0; i < len(nums); i++ {
		rui += nums[i]
		ans = max(ans, rui-ruiMin)
		if rui < ruiMin {
			ruiMin = rui
		}
	}
	return ans
}

func maxSubArray0(nums []int) int {
	rui := make([]int, len(nums))
	rui[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		rui[i] = rui[i-1] + nums[i]
	}

	ans, minv := rui[0], 0
	for i := range rui {
		ans = max(ans, rui[i]-minv)
		if rui[i] < minv {
			minv = rui[i]
		}
	}

	return ans
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) == 6)
	fmt.Println(maxSubArray([]int{-1}) == -1)
}
