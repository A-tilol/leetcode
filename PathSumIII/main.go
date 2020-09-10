package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func dfs(node *TreeNode, sum int) []int {
	if node.Left == nil && node.Right == nil {
		if node.Val == sum {
			ans++
		}
		return []int{node.Val}
	}

	lres, rres := make([]int, 0), make([]int, 0)
	if node.Left != nil {
		lres = dfs(node.Left, sum)
	}
	if node.Right != nil {
		rres = dfs(node.Right, sum)
	}

	lres = append(lres, rres...)
	for i := range lres {
		lres[i] += node.Val
		if lres[i] == sum {
			ans++
		}
	}
	if node.Val == sum {
		ans++
	}
	lres = append(lres, node.Val)

	return lres
}

func pathSum(root *TreeNode, sum int) int {
	ans = 0
	if root == nil {
		return ans
	}
	dfs(root, sum)
	return ans
}

func dfs1(node *TreeNode, sum int) map[int]int {
	if node.Left == nil && node.Right == nil {
		if node.Val == sum {
			ans++
			// fmt.Println(node.Val)
		}
		return map[int]int{node.Val: 1}
	}

	lres, rres := make(map[int]int), make(map[int]int)
	if node.Left != nil {
		lres = dfs1(node.Left, sum)
	}
	if node.Right != nil {
		rres = dfs1(node.Right, sum)
	}

	ret := make(map[int]int)
	for k, v := range lres {
		ret[k+node.Val] += v
	}
	for k, v := range rres {
		ret[k+node.Val] += v
	}
	ret[node.Val]++

	if v, ok := ret[sum]; ok {
		ans += v
		// fmt.Println(node.Val)
	}

	return ret
}

func main() {
}
