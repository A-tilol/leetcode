package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans [][]int

func dfs(node *TreeNode, depth int) {
	if node == nil {
		return
	}

	if len(ans)-1 < depth {
		ans = append(ans, []int{})
	}
	ans[depth] = append(ans[depth], node.Val)

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
func levelOrder(root *TreeNode) [][]int {
	ans = make([][]int, 0)
	dfs(root, 0)
	return ans
}

func main() {
}
