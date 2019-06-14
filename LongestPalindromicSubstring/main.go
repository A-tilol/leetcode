package longestpalindromicsubstring

func longestPalindrome(s string) string {
	n := len(s)
	b := make([]byte, 0, n)
	for i := n - 1; i >= 0; i-- {
		b = append(b, s[i])
	}
	ss := string(b)

	ans := ""
	anslen := 0
	for i := 0; i < n; i++ {
		for j := n; j > i; j-- {
			if j-i <= anslen {
				break
			}
			if s[i:j] == ss[n-j:n-i] {
				ans = s[i:j]
				anslen = j - i
				break
			}
		}
	}
	return ans
}

func longestPalindrome1(s string) string {
	n := len(s)
	ss := ""
	for i := n - 1; i >= 0; i-- {
		ss += string(s[i])
	}

	ans := ""
	anslen := 0
	for i := 0; i < n; i++ {
		for j := n; j > i; j-- {
			if j-i <= anslen {
				break
			}
			if s[i:j] == ss[n-j:n-i] {
				ans = s[i:j]
				anslen = j - i
				break
			}
		}
	}
	return ans
}
