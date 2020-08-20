package main

import (
	"fmt"
)

var s string
var ans []string

func dfs(i int, letter string) {
	if i == len(s) {
		ans = append(ans, letter)
		return
	}

	num := int(s[i]) - 50
	l := 'a' + num*3
	if num > 5 {
		l++
	}
	n := 3
	if num == 5 || num == 7 {
		n = 4
	}
	for j := 0; j < n; j++ {
		dfs(i+1, letter+string(l+j))
	}
}

func letterCombinations(digits string) []string {
	ans = make([]string, 0)
	s = digits
	if len(s) == 0 {
		return ans
	}
	dfs(0, "")
	return ans
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("7"))
	fmt.Println(letterCombinations("9"))
}
