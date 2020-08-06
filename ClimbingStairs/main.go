package main

import (
	"fmt"
)

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	pre0, pre1 := 1, 1
	for i := 2; i < n; i++ {
		pre0, pre1 = pre1, pre0+pre1
	}
	return pre0 + pre1
}

func main() {
	fmt.Println(climbStairs(1) == 1)
	fmt.Println(climbStairs(2) == 2)
	fmt.Println(climbStairs(3) == 3)
	fmt.Println(climbStairs(5) == 8)
}
