package containerwithmostwater

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	var i, ans int
	j := len(height) - 1
	for i < j {
		ans = max(ans, (j-i)*min(height[i], height[j]))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return ans
}

func maxArea1(height []int) int {
	n := len(height)
	var i, j, ans, v int
	for i = 0; i < n; i++ {
		for j = i + 1; j < n; j++ {
			v = (j - i) * min(height[i], height[j])
			ans = max(ans, v)
		}
	}
	return ans
}
