package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(merged, t1, t2 *TreeNode) {
	var t1l, t1r, t2l, t2r *TreeNode = nil, nil, nil, nil
	if t1 != nil {
		merged.Val += t1.Val
		t1l, t1r = t1.Left, t1.Right
	}
	if t2 != nil {
		merged.Val += t2.Val
		t2l, t2r = t2.Left, t2.Right
	}

	if t1l != nil || t2l != nil {
		merged.Left = &TreeNode{0, nil, nil}
		dfs(merged.Left, t1l, t2l)
	}
	if t1r != nil || t2r != nil {
		merged.Right = &TreeNode{0, nil, nil}
		dfs(merged.Right, t1r, t2r)
	}
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	merged := &TreeNode{}
	dfs(merged, t1, t2)

	return merged
}

func main() {

}
