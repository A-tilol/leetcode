package main

import (
	"fmt"
	"reflect"
)

func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		for nums[i] != i+1 && nums[i] != -1 {
			if nums[i]-1 < i || nums[nums[i]-1] == nums[i] {
				nums[i], nums[nums[i]-1] = -1, nums[i]
			} else {
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			}
		}
	}

	ans := make([]int, 0)
	for i := range nums {
		if nums[i] == -1 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func main() {
	fmt.Println(reflect.DeepEqual(
		findDisappearedNumbers(
			[]int{4, 3, 2, 7, 8, 2, 3, 1},
		), []int{5, 6}))
	fmt.Println(reflect.DeepEqual(
		findDisappearedNumbers(
			[]int{1, 2, 3, 4, 5},
		), []int{}))
	fmt.Println(reflect.DeepEqual(
		findDisappearedNumbers(
			[]int{1, 1, 2, 3, 4, 5},
		), []int{6}))
	fmt.Println(reflect.DeepEqual(
		findDisappearedNumbers(
			[]int{1, 1, 1},
		), []int{2, 3}))
	fmt.Println(reflect.DeepEqual(
		findDisappearedNumbers(
			[]int{2, 2, 2},
		), []int{1, 3}))
}
