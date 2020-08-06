package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// less space complexity
func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
	n := len(cost)
	for i := 2; i < n; i++ {
		cost[i] += min(cost[i-2], cost[i-1])
	}
	return cost[n-1]
}

func minCostClimbingStairs0(cost []int) int {
	n := len(cost)
	totalCost := make([]int, n+1)
	for i := 2; i <= n; i++ {
		totalCost[i] = min(totalCost[i-2]+cost[i-2], totalCost[i-1]+cost[i-1])
	}
	return totalCost[n]
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
