package main

import (
	"fmt"
	"strconv"
)

var cToNum map[byte]int

func init() {
	cToNum = make(map[byte]int)
	for i := 0; i < 9; i++ {
		cToNum[byte('1'+i)] = i + 1
	}
}

func dfs(s string) []byte {
	ret := make([]byte, 0)
	i, times := 0, 0
	for i < len(s) {
		if _, ok := cToNum[s[i]]; ok {
			numEnd := i + 1
			for s[numEnd] != '[' {
				numEnd++
			}
			times, _ = strconv.Atoi(s[i:numEnd])
			i = numEnd + 1
			break
		}
		ret = append(ret, s[i])
		i++
	}

	if times == 0 {
		return ret
	}

	j, pcnt := i, 1
	for pcnt > 0 {
		if s[j] == '[' {
			pcnt++
		} else if s[j] == ']' {
			pcnt--
		}
		j++
	}

	res := dfs(s[i : j-1])
	for k := 0; k < times; k++ {
		ret = append(ret, res...)
	}

	return append(ret, dfs(s[j:])...)
}

func decodeString(s string) string {
	return string(dfs(s))
}

func main() {
	fmt.Println(decodeString("3[a]2[bc]") == "aaabcbc")
	fmt.Println(decodeString("3[a2[c]]") == "accaccacc")
	fmt.Println(decodeString("2[abc]3[cd]ef") == "abcabccdcdcdef")
	fmt.Println(decodeString("abc3[cd]xyz") == "abccdcdcdxyz")
	fmt.Println(decodeString("10[ab]") == "abababababababababab")
}
