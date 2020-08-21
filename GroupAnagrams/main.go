package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		bs := []byte(s)
		sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
		t := string(bs)
		if _, ok := m[t]; !ok {
			m[t] = make([]string, 0)
		}
		m[t] = append(m[t], s)

	}
	ans := make([][]string, 0, len(m))
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	fmt.Println(groupAnagrams(
		[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
	))
}
