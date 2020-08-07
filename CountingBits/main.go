package main

import (
	"fmt"
)

func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}

	dp := make([]int, num+1)
	dp[0], dp[1] = 0, 1

	rui2, cnt := 2, 0
	for i := 2; i <= num; i++ {
		if i == rui2 {
			rui2 <<= 1
			cnt = 0
		}
		dp[i] = dp[cnt] + 1
		cnt++
	}
	return dp
}

func main() {
	fmt.Println(countBits(0))
	fmt.Println(countBits(1))
	fmt.Println(countBits(2))
	fmt.Println(countBits(15))
}
