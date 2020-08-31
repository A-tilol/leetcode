package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(source, dest *TreeNode) {
	if source == nil {
		return
	}

	dest.Val = source.Val
	if source.Right != nil {
		dest.Left = &TreeNode{}
	}
	if source.Left != nil {
		dest.Right = &TreeNode{}
	}

	dfs(source.Left, dest.Right)
	dfs(source.Right, dest.Left)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	dest := &TreeNode{}
	dfs(root, dest)
	return dest
}

func main() {
}
