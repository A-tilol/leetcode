package main

import (
	"fmt"
)

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		if cur-1 == i {
			continue
		}

		if cur == nums[cur-1] {
			return cur
		}

		nums[i], nums[cur-1] = nums[cur-1], cur
		i--
	}
	return 0
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Println(findDuplicate([]int{1, 1}))
	fmt.Println(findDuplicate([]int{1, 1, 2}))
	fmt.Println(findDuplicate([]int{2, 6, 4, 1, 3, 1, 5}))
}
