package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	maxi := 0
	for i := range nums {
		if i > maxi {
			return false
		}
		if i+nums[i] > maxi {
			maxi = i + nums[i]
			if maxi >= len(nums) {
				return true
			}
		}
	}
	return true
}

func main() {
	fmt.Println(canJump(
		[]int{2, 3, 1, 1, 4},
	))
	fmt.Println(canJump(
		[]int{3, 2, 1, 0, 4},
	))
}
