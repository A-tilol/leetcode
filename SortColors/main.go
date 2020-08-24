package main

func sortColors(nums []int) {
	i, zero, two := 0, 0, len(nums)-1
	for i < len(nums) {
		if nums[i] == 0 && i > zero {
			nums[zero], nums[i] = nums[i], nums[zero]
			zero++
			continue
		}
		if nums[i] == 2 && i < two {
			nums[two], nums[i] = nums[i], nums[two]
			two--
			continue
		}
		i++
	}
	// fmt.Println(nums)
}

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0})
	sortColors([]int{1, 2, 0})
}
