package main

import (
	"fmt"
)

func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	for i := range s {
		for j := 0; i-j >= 0 && i+j+1 < n; j++ {
			if s[i-j] != s[i+j+1] {
				break
			}
			ans++
		}
		for j := 0; i-j >= 0 && i+j < n; j++ {
			if s[i-j] != s[i+j] {
				break
			}
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("aaa"))
	fmt.Println(countSubstrings("abcbade"))
	fmt.Println(countSubstrings("aaaaa"))
}
