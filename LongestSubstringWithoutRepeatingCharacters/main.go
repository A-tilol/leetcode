package longestsubstringwithoutrepeatingcharacters

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func lengthOfLongestSubstring(s string) int {
	var i, start, ans = 0, 0, 0
	m := make(map[byte]int, 256)
	for i = 0; i < len(s); i++ {
		if j, ok := m[s[i]]; ok && j >= start {
			ans = max(ans, i-start)
			start = j + 1
		}
		m[s[i]] = i
	}
	return max(ans, i-start)
}

func lengthOfLongestSubstring2(s string) int {
	var i, start, ans = 0, 0, 0
	m := make(map[byte]struct{})
	for i = 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			ans = max(ans, i-start)
			delete(m, s[start])
			start++
			i--
			continue
		}
		m[s[i]] = struct{}{}
	}
	return max(ans, i-start)
}

func lengthOfLongestSubstring1(s string) int {
	ans := 0
	j := 0
	for i := range s {
		m := make(map[byte]struct{})
		for j = i; j < len(s); j++ {
			if _, ok := m[s[j]]; ok {
				break
			}
			m[s[j]] = struct{}{}
		}
		ans = max(ans, j-i)
	}
	return ans
}
