package main

func moveZeroes(nums []int) {
	zeroCnt := 0
	for i := range nums {
		if nums[i] == 0 {
			zeroCnt++
			continue
		}
		nums[i-zeroCnt], nums[i] = nums[i], nums[i-zeroCnt]
	}
}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{1, 2, 0, 0})
	moveZeroes([]int{0, 0, 0})
	moveZeroes([]int{1})
	moveZeroes([]int{0})
}
