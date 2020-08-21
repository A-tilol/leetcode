package main

func rotate(matrix [][]int) {
	l, r := 0, len(matrix)-1
	for l < r {
		for i := 0; i < r-l; i++ {
			matrix[l][l+i], matrix[l+i][r], matrix[r][r-i], matrix[r-i][l] =
				matrix[r-i][l], matrix[l][l+i], matrix[l+i][r], matrix[r][r-i]
		}
		l++
		r--
	}
}

func main() {
	rotate([][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	})

	rotate([][]int{
		[]int{5, 1, 9, 11},
		[]int{2, 4, 8, 10},
		[]int{13, 3, 6, 7},
		[]int{15, 14, 12, 16},
	})
}
