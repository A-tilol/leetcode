package zigzagconversion

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	a := make([]byte, 0, len(s))
	f := 2 * (numRows - 1)
	r := 0
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(s); j++ {
			r = j % f
			if min(r, f-r) == i {
				a = append(a, s[j])
			}
		}
	}
	return string(a)
}

func convert2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	n := len(s)
	a := make([][]byte, numRows)
	ans := make([]byte, 0, n)
	f := 2 * (numRows - 1)
	m := n / f
	var i, r int

	for i = 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			a[i] = make([]byte, 0, m)
		}
		a[i] = make([]byte, 0, 2*m)
	}

	for i = 0; i < n; i++ {
		r = i % f
		r = min(r, f-r)
		a[r] = append(a[r], s[i])
	}

	for i = 0; i < numRows; i++ {
		ans = append(ans, a[i]...)
	}
	return string(ans)
}

func convert1(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	a := make([][]byte, numRows)
	f := 2 * (numRows - 1)
	r := 0
	for i := 0; i < len(s); i++ {
		r = i % f
		r = min(r, f-r)
		a[r] = append(a[r], s[i])
	}

	ans := make([]byte, 0, len(s))
	for i := 0; i < numRows; i++ {
		ans = append(ans, a[i]...)
	}
	return string(ans)
}
