package main

import (
	"fmt"
	"math"
)

func dailyTemperatures(T []int) []int {
	stack := make([]int, 0)
	ans := make([]int, len(T))
	for i := range T {
		for len(stack) > 0 && T[stack[len(stack)-1]] < T[i] {
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dailyTemperatures1(T []int) []int {
	tmp := make([][]int, 101)
	for i := range tmp {
		tmp[i] = make([]int, 0)
	}

	ans := make([]int, len(T))
	minT := math.MaxInt32
	for i := range T {
		for j := minT; j < T[i]; j++ {
			for _, idx := range tmp[j] {
				ans[idx] = i - idx
			}
			if len(tmp[j]) > 0 {
				tmp[j] = make([]int, 0)
			}
		}
		tmp[T[i]] = append(tmp[T[i]], i)
		minT = min(minT, T[i])
	}

	return ans
}

func main() {
	fmt.Println(dailyTemperatures([]int{
		73, 74, 75, 71, 69, 72, 76, 73,
	}))
}
