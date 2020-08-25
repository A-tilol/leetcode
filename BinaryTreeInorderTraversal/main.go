package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// repeatively -----------------------------------
func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		if cur == nil {
			stack = stack[:len(stack)-1]
			continue
		}

		if cur.Left != nil {
			stack = append(stack, cur.Left)
			cur.Left = nil
			continue
		}

		ans = append(ans, cur.Val)
		stack = stack[:len(stack)-1]

		stack = append(stack, cur.Right)
	}

	return ans
}

// dfs -----------------------------------------
var ans []int

func dfs(node *TreeNode) {
	if node == nil {
		// fmt.Println(nil)
		return
	}
	fmt.Println(node.Val)

	dfs(node.Left)
	ans = append(ans, node.Val)
	dfs(node.Right)
}

func inorderTraversal1(root *TreeNode) []int {
	ans = make([]int, 0)
	dfs(root)
	return ans
}

func main() {
}
