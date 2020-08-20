package main

import "fmt"

var ans []string

func dfs(s string, l, r int) {
	if l == 0 && r == 0 {
		ans = append(ans, s)
		return
	}

	if l > 0 {
		dfs(s+"(", l-1, r)
	}

	if r > l {
		dfs(s+")", l, r-1)
	}
}

func generateParenthesis(n int) []string {
	ans = make([]string, 0)
	if n == 0 {
		return ans
	}
	dfs("", n, n)
	return ans
}

func main() {
	fmt.Println(generateParenthesis(3))
}
