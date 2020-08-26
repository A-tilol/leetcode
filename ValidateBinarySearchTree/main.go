package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}
	if node.Val <= lower || node.Val >= upper {
		return false
	}

	if !dfs(node.Left, lower, node.Val) {
		return false
	}
	if !dfs(node.Right, node.Val, upper) {
		return false
	}

	return true
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt64, math.MaxInt64)
}

func main() {
}
