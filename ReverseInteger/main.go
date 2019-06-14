package reverseinteger

import (
	"math"
	"strconv"
)

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func reverse(x int) int {
	s := strconv.Itoa(abs(x))
	n := len(s)
	b := make([]byte, 0, n+1)
	if x < 0 {
		b = append(b, byte('-'))
	}
	for i := range s {
		b = append(b, s[n-i-1])
	}
	ans, _ := strconv.Atoi(string(b))
	if ans <= math.MaxInt32 && ans >= math.MinInt32 {
		return ans
	}
	return 0
}
