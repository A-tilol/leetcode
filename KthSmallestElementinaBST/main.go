package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans, ith, kk int

func dfs(node *TreeNode) {
	if ith == kk {
		return
	}
	if node == nil {
		return
	}

	dfs(node.Left)

	ith++
	if ith == kk {
		ans = node.Val
		return
	}

	dfs(node.Right)

	return
}

func kthSmallest(root *TreeNode, k int) int {
	ith, kk = 0, k
	dfs(root)
	return ans
}

func main() {
}
