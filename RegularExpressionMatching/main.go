package regularexpressionmatching

import "regexp"

func isMatch(s string, p string) bool {
	ans, _ := regexp.MatchString("^"+p+"$", s)
	return ans
}
