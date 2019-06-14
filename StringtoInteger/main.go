package stringtointeger

import (
	"math"
	"strconv"
	"strings"
)

var valid = "0123456789"

func removePrefixWhiteSpace(s string) string {
	for i := range s {
		if s[i] != byte(' ') {
			return s[i:]
		}
	}
	return ""
}

func extractValidSub(s string) string {
	for i := range s {
		if !strings.Contains(valid, string(s[i])) {
			return s[:i]
		}
	}
	return s
}

func myAtoi(str string) int {
	s := removePrefixWhiteSpace(str)

	if s == "" || s == "-" || s == "+" {
		return 0
	}

	negative := false
	if s[0] == byte('-') {
		negative = true
		s = s[1:]
	} else if s[0] == byte('+') {
		s = s[1:]
	}

	if !strings.Contains(valid, string(s[0])) {
		return 0
	}

	s = extractValidSub(s)

	ans, _ := strconv.Atoi(s)
	if negative {
		ans = -ans
	}
	if ans < math.MinInt32 {
		return math.MinInt32
	}
	if ans > math.MaxInt32 {
		return math.MaxInt32
	}
	return ans
}
