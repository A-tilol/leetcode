package main

import (
	"fmt"
	"sort"
)

// less space and time complexity
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	ans := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r] == -nums[i] {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				for l < len(nums)-1 && nums[l] == nums[l+1] {
					l++
				}
				for r > i+1 && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if nums[l]+nums[r] > -nums[i] {
				r--
			} else {
				l++
			}
		}
	}
	return ans
}

// add early break
func threeSum3(nums []int) [][]int {
	sort.Ints(nums)

	numMap := make(map[int]int)
	for i := range nums {
		numMap[nums[i]] = i
	}

	ans := make([][]int, 0)
	already := make(map[int]map[int]struct{})
	for i := range nums {
		if nums[i] > 0 {
			break
		}

		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] > 0 {
				break
			}

			want := -(nums[i] + nums[j])
			if idx, ok := numMap[want]; !ok || idx <= j {
				continue
			}

			if _, ok := already[nums[i]][nums[j]]; ok {
				continue
			}

			// fmt.Println(nums[i], nums[j], nums[k])
			ans = append(ans, []int{nums[i], nums[j], want})

			if _, ok := already[nums[i]]; !ok {
				already[nums[i]] = make(map[int]struct{})
			}
			already[nums[i]][nums[j]] = struct{}{}
		}
	}
	return ans
}

// less time complexity. don't generate string
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)

	numMap := make(map[int]int)
	for i := range nums {
		numMap[nums[i]] = i
	}

	ans := make([][]int, 0)
	already := make(map[int]map[int]struct{})
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			want := -(nums[i] + nums[j])
			if idx, ok := numMap[want]; !ok || idx <= j {
				continue
			}

			if _, ok := already[nums[i]][nums[j]]; ok {
				continue
			}

			// fmt.Println(nums[i], nums[j], nums[k])
			ans = append(ans, []int{nums[i], nums[j], want})

			if _, ok := already[nums[i]]; !ok {
				already[nums[i]] = make(map[int]struct{})
			}
			already[nums[i]][nums[j]] = struct{}{}
		}
	}
	return ans
}

// less time complexity. O(n**2)
func threeSum1(nums []int) [][]int {
	sort.Ints(nums)

	numMap := make(map[int]int)
	for i := range nums {
		numMap[nums[i]] = i
	}

	ans := make([][]int, 0)
	triplets := make(map[string]struct{})
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			want := -(nums[i] + nums[j])
			if idx, ok := numMap[want]; !ok || idx <= j {
				continue
			}

			key := fmt.Sprintf("%d,%d,%d", nums[i], nums[j], want)
			if _, ok := triplets[key]; ok {
				continue
			}

			// fmt.Println(nums[i], nums[j], nums[k])
			ans = append(ans, []int{nums[i], nums[j], want})
			triplets[key] = struct{}{}
		}
	}
	return ans
}

// Failed: Time Limit Exceeded. O(n**3)
func threeSum0(nums []int) [][]int {
	sort.Ints(nums)

	ans := make([][]int, 0)
	triplets := make(map[string]struct{})
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] != 0 {
					continue
				}

				key := fmt.Sprintf("%d,%d,%d", nums[i], nums[j], nums[k])
				if _, ok := triplets[key]; ok {
					continue
				}

				// fmt.Println(nums[i], nums[j], nums[k])
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				triplets[key] = struct{}{}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-1, 0, 0, 0, 1, 1}))
}
