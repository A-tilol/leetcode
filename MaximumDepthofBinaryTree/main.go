package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(node *TreeNode, depth int) {
	if node == nil {
		return
	}
	ans = max(ans, depth)
	dfs(node.Left, depth+1)
	dfs(node.Right, depth+1)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	ans = 0
	dfs(root, 1)
	return ans
}

func main() {
}
