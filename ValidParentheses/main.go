package main

import (
	"fmt"
)

var parenthesesMap = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

func isValid(s string) bool {
	// fmt.Printf("\n\n%v\n", s)
	stack := make([]rune, 0)
	for _, c := range s {
		// fmt.Println(c, stack, parenthesesMap)
		if closeParenthese, ok := parenthesesMap[c]; ok {
			stack = append(stack, closeParenthese)
			continue
		}
		if len(stack) == 0 || c != stack[len(stack)-1] {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isValid("()") == true)
	fmt.Println(isValid("()[]{}") == true)
	fmt.Println(isValid("(]") == false)
	fmt.Println(isValid("([)]") == false)
	fmt.Println(isValid("{[]}") == true)
	fmt.Println(isValid("(){[()]}") == true)
	fmt.Println(isValid("(){[(])}") == false)
}
