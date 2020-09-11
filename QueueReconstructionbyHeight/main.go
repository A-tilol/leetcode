package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool { return people[i][0] < people[j][0] })
	fmt.Println(people)

	used := make([]bool, len(people))
	nums := make(map[int]int)
	for _, person := range people {
		nums[person[0]] = 0
	}

	done := 0
	ans := make([][]int, 0, len(people))
	for range people {
		// fmt.Println(ans)
		// fmt.Println(used)
		// fmt.Println(nums)
		// fmt.Println()
		frontNum, preh := done, people[0][0]
		for i, person := range people {
			h, k := person[0], person[1]
			if h > preh {
				frontNum -= nums[preh]
				preh = h
			}
			if used[i] || k != frontNum {
				continue
			}

			ans = append(ans, person)
			used[i] = true
			done++
			nums[h]++
			break
		}
	}

	return ans
}

func main() {
	fmt.Println(reconstructQueue([][]int{
		[]int{7, 0}, []int{4, 4}, []int{7, 1}, []int{5, 0}, []int{6, 1}, []int{5, 2},
	}))
}
