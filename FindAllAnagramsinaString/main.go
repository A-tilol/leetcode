package main

import (
	"fmt"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// use slice. 5 times faster than the map's implementaion
func findAnagrams(s string, p string) []int {
	ans := make([]int, 0)
	if s == "" || len(s) < len(p) {
		return ans
	}

	cnt := make([]int, 27)
	for i := range p {
		cnt[p[i]-'a']++
	}

	diff := len(p)
	for i := range s {
		if i-len(p) >= 0 {
			c := s[i-len(p)] - 'a'
			if abs(cnt[c]+1) < abs(cnt[c]) {
				diff--
			} else {
				diff++
			}
			cnt[c]++
		}

		c := s[i] - 'a'
		if abs(cnt[c]-1) < abs(cnt[c]) {
			diff--
		} else {
			diff++
		}
		cnt[c]--

		if diff == 0 {
			ans = append(ans, i-len(p)+1)
		}
	}

	return ans
}

// use map
func findAnagrams1(s string, p string) []int {
	ans := make([]int, 0)
	if s == "" || len(s) < len(p) {
		return ans
	}

	cnt := make(map[byte]int)
	for i := 0; i < 27; i++ {
		cnt[byte('a'+i)] = 0
	}
	for i := range p {
		cnt[p[i]]++
	}

	diff := len(p)
	for i := range s {
		if i-len(p) >= 0 {
			if abs(cnt[s[i-len(p)]]+1) < abs(cnt[s[i-len(p)]]) {
				diff--
			} else {
				diff++
			}
			cnt[s[i-len(p)]]++
		}

		if abs(cnt[s[i]]-1) < abs(cnt[s[i]]) {
			diff--
		} else {
			diff++
		}
		cnt[s[i]]--

		if diff == 0 {
			ans = append(ans, i-len(p)+1)
		}
	}

	return ans
}

func main() {
	fmt.Println(findAnagrams("", "a"))
	fmt.Println(findAnagrams("a", "aaa"))
	fmt.Println(findAnagrams("aa", "aa"))
	fmt.Println(findAnagrams("aaaa", "a"))
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}
