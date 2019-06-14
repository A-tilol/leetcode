package PalindromeNumber

import "strconv"

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	n := len(s)
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[n-1-i] = s[i]
	}
	return s == string(b)
}
