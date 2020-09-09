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

func dfs(node *TreeNode) int {
	if node.Left == nil && node.Right == nil {
		return 1
	}

	ldepth, rdepth := 0, 0
	if node.Left != nil {
		ldepth = dfs(node.Left)
	}
	if node.Right != nil {
		rdepth = dfs(node.Right)
	}

	ans = max(ans, ldepth+rdepth)

	return max(ldepth, rdepth) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans = 0
	if root == nil {
		return ans
	}
	dfs(root)
	return ans
}

func main() {
}
