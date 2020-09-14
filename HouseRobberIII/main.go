package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var memo []map[*TreeNode]int

func dfs(node *TreeNode, hasRedParent int) int {
	if v, ok := memo[hasRedParent][node]; ok {
		return v
	}
	if node == nil {
		return 0
	}

	ret1, ret2 := 0, 0
	ret1 = dfs(node.Left, 0) + dfs(node.Right, 0)
	if hasRedParent == 0 {
		ret2 = dfs(node.Left, 1) + dfs(node.Right, 1) + node.Val
	}

	memo[hasRedParent][node] = max(ret1, ret2)
	return memo[hasRedParent][node]
}

func rob(root *TreeNode) int {
	memo = []map[*TreeNode]int{map[*TreeNode]int{}, map[*TreeNode]int{}}
	return dfs(root, 0)
}

func main() {
}
