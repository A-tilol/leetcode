package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var lefts, rights []int

func dfsLeft(node *TreeNode) {
	if node == nil {
		lefts = append(lefts, -1)
		return
	}
	lefts = append(lefts, node.Val)
	dfsLeft(node.Left)
	dfsLeft(node.Right)

}

func dfsRight(node *TreeNode) {
	if node == nil {
		rights = append(rights, -1)
		return
	}
	rights = append(rights, node.Val)
	dfsRight(node.Right)
	dfsRight(node.Left)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lefts, rights = make([]int, 0), make([]int, 0)
	dfsLeft(root.Left)
	dfsRight(root.Right)
	if len(lefts) != len(rights) {
		return false
	}
	for i := range lefts {
		if lefts[i] != rights[i] {
			return false
		}
	}
	return true
}

func main() {
}
